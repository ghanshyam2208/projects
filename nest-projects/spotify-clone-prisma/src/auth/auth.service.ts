import {
  Injectable,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';
import { UsersRepository } from '../users/users.repository';
import { LoginPayload } from '../users/users.validations';
import { CryptoHelper } from 'src/common/crypto.helper';
import { JwtService } from '@nestjs/jwt';

@Injectable()
export class AuthService {
  constructor(
    private readonly usersRepository: UsersRepository,
    private readonly cryptoHelper: CryptoHelper,
    private jwtService: JwtService,
  ) {}

  async loginUser(
    loginPayload: LoginPayload,
  ): Promise<{ accessToken: string }> {
    const user = await this.usersRepository.findUserByEmail(loginPayload.email);
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
      accessToken: await this.jwtService.signAsync(payload, {
        secret: process.env.CRYPTO_SECRET,
        expiresIn: '1d',
      }),
    };
  }
}
