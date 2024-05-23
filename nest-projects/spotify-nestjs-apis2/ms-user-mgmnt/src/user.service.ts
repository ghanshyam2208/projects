import { BadRequestException, Injectable } from '@nestjs/common';
import { CreateUserPayload, sanitizedUserResponse } from './user.validation';
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
}
