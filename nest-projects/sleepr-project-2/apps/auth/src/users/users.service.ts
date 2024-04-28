import { Injectable, UnauthorizedException } from '@nestjs/common';
import * as bcrypt from 'bcryptjs';
import { UsersRepository } from './users.repository';
import { CreateUserPayload } from './users.validation';

@Injectable()
export class UsersService {
  constructor(private readonly usersRepository: UsersRepository) {}

  async createUser(createUserPayload: CreateUserPayload) {
    this.usersRepository.create({
      ...createUserPayload,
      password: await bcrypt.hash(createUserPayload.password, 10),
    });
  }

  async validateUser(email: string, password: string) {
    const user = await this.usersRepository.findOne({
      email,
    });

    if (!user) {
      throw new UnauthorizedException('Username or password is not valid');
    }

    if (!(await bcrypt.compare(password, user.password))) {
      throw new UnauthorizedException('Username or password is not valid');
    }
    return user;
  }
}
