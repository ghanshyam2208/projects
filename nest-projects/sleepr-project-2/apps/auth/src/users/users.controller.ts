import { Body, Controller, Post } from '@nestjs/common';
import {
  CreateUserPayload,
  CreateUsersValidationPipe,
} from './users.validation';
import { UsersService } from './users.service';

@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post()
  async createUser(
    @Body(new CreateUsersValidationPipe()) createUserPayload: CreateUserPayload,
  ) {
    return this.usersService.createUser(createUserPayload);
  }
}
