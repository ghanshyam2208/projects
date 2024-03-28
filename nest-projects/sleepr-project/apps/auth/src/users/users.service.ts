import { Injectable } from '@nestjs/common';
import { UsersRepository } from './users.repository';
import { CreateUsersDto } from './users.validations';

@Injectable()
export class UsersService {
  constructor(private readonly usersRepository: UsersRepository) {}

  async createUser(createUsersDto: CreateUsersDto) {
    return await this.usersRepository.create(createUsersDto);
  }
}
