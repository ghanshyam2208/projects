import { Injectable } from '@nestjs/common';
import { UsersRepository } from './users.repository';
import { CreateUserPayload } from './users.validation';

@Injectable()
export class UsersService {
  constructor(private readonly usersRepository: UsersRepository) {}

  async createUser(createUserPayload: CreateUserPayload) {
    this.usersRepository.create(createUserPayload);
  }
}
