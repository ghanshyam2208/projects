import { Body, Controller, Post } from '@nestjs/common';
import {
  CreateUserPayload,
  CreateUserValidationPipe,
} from '../users/users.validations';
import { UsersService } from '../users/users.service';

@Controller('auth')
export class AuthController {
  constructor(private readonly userService: UsersService) {}

  @Post()
  signUp(
    @Body(new CreateUserValidationPipe()) createUserPayload: CreateUserPayload,
  ) {
    return this.userService.createUser(createUserPayload);
  }
}
