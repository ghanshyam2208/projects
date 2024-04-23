import { Injectable, InternalServerErrorException } from '@nestjs/common';
import { UsersRepository } from './users.repository';
import { CreateUserPayload, sanitizeUserResponse } from './users.validations';
import { CryptoHelper } from '../common/crypto.helper';
import { Users } from '@prisma/client';

@Injectable()
export class UsersService {
  constructor(
    private readonly usersRepository: UsersRepository,
    private readonly cryptoHelper: CryptoHelper,
  ) {}

  async createUser(createUserPayload: CreateUserPayload): Promise<Users> {
    try {
      createUserPayload.password = this.cryptoHelper.encrypt(
        createUserPayload.password,
      );
      const user = await this.usersRepository.createUser(createUserPayload);
      return sanitizeUserResponse(user);
    } catch (error) {
      console.error(error);
      throw new InternalServerErrorException(error);
    }
  }
}
