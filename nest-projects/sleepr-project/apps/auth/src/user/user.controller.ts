import { Body, Controller, Get, Post } from '@nestjs/common';
import {
  CreateUserPayload,
  CreateUserPayloadValidationPipe,
} from './dto/user.payload';
import { UserService } from './user.service';

@Controller('users')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Get()
  getHello(): string {
    return 'return some code';
  }

  @Post()
  createUser(
    @Body(new CreateUserPayloadValidationPipe())
    createUserPayload: CreateUserPayload,
  ) {
    return this.userService.createUser(createUserPayload);
  }
}
