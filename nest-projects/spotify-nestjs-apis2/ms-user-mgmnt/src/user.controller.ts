import { Body, Controller, Post } from '@nestjs/common';
import { CreateUserPayload, CreateUserValidationPipe } from './user.validation';
import { UserService } from './user.service';

@Controller('user')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Post()
  registerUser(
    @Body(new CreateUserValidationPipe()) createUserPayload: CreateUserPayload,
  ) {
    return this.userService.registerUser(createUserPayload);
  }
}
