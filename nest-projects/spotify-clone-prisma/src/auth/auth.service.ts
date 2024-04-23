import {
  BadRequestException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { UsersRepository } from '../users/user.repository';
import { LoginPayload, sanitizeUserResponse } from '../users/users.validations';
import { CryptoHelper } from 'src/common/crypto.helper';

@Injectable()
export class AuthService {
  constructor(
    private readonly usersRepository: UsersRepository,
    private readonly cryptoHelper: CryptoHelper,
  ) {}

  async loginUser(loginPayload: LoginPayload) {
    const user = await this.usersRepository.findUserByEmail(loginPayload.email);
    if (!user) {
      throw new NotFoundException(
        `User not found with email ${loginPayload.email}`,
      );
    }
    const decryptedPassword = this.cryptoHelper.decrypt(user.password);
    if (decryptedPassword !== loginPayload.password) {
      throw new BadRequestException(`Username or password is wrong`);
    }
    return sanitizeUserResponse(user);
  }
}
