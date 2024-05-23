import {
  BadRequestException,
  Injectable,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';
import {
  CreateUserPayload,
  LoginPayload,
  sanitizedUserResponse,
} from './user.validation';
import { UserRepository } from './user.repository';
import { CryptoHelper } from './prisma/crypto.helper';

@Injectable()
export class UserService {
  constructor(
    private readonly userRepository: UserRepository,
    private readonly cryptoHelper: CryptoHelper,
  ) {}

  async registerUser(createUserPayload: CreateUserPayload) {
    try {
      createUserPayload.password = this.cryptoHelper.encrypt(
        createUserPayload.password,
      );
      const user = await this.userRepository.registerUser(createUserPayload);
      return sanitizedUserResponse(user);
    } catch (error) {
      if (error.code === 'P2002') {
        throw new BadRequestException('User email/username already exists');
      }
      throw error;
    }
  }

  async loginUser(loginPayload: LoginPayload) {
    const user = await this.userRepository.findUserByEmail(loginPayload.email);
    if (!user) {
      throw new NotFoundException(
        `User not found with email ${loginPayload.email}`,
      );
    }
    const decryptedPassword = this.cryptoHelper.decrypt(user.password);
    if (decryptedPassword !== loginPayload.password) {
      throw new UnauthorizedException(`Username or password is wrong`);
    }
    const payload = { sub: user.id, email: user.email };
    return {
      // accessToken: await this.jwtService.signAsync(payload, {
      //   secret: process.env.CRYPTO_SECRET,
      //   expiresIn: '1d',
      // }),
      login: true,
      payload: {
        ...payload,
      },
    };
  }
}
