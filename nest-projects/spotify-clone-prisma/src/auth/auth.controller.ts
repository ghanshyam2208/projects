import { Body, Controller, Post } from '@nestjs/common';
import {
  CreateUserPayload,
  CreateUserValidationPipe,
  LoginPayload,
  LoginValidationPipe,
} from '../users/users.validations';
import { UsersService } from '../users/users.service';
import { AuthService } from './auth.service';

@Controller('auth')
export class AuthController {
  constructor(
    private readonly userService: UsersService,
    private readonly authService: AuthService,
  ) {}

  @Post('signup')
  signUp(
    @Body(new CreateUserValidationPipe()) createUserPayload: CreateUserPayload,
  ) {
    return this.userService.createUser(createUserPayload);
  }

  @Post('login')
  login(@Body(new LoginValidationPipe()) loginPayload: LoginPayload) {
    return this.authService.loginUser(loginPayload);
  }
}
