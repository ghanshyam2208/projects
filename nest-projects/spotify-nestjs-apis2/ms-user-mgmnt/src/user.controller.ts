import { Body, Controller, Get, Post, UseGuards } from '@nestjs/common';
import {
  CreateUserPayload,
  CreateUserValidationPipe,
  LoginPayload,
  LoginValidationPipe,
} from './user.validation';
import { UserService } from './user.service';
import { PostTodoDTO } from 'proto/todo';
import { AuthGuard } from './guards/user.guard';

@Controller('user')
export class UserController {
  constructor(private readonly userService: UserService) {}

  @Post()
  registerUser(
    @Body(new CreateUserValidationPipe()) createUserPayload: CreateUserPayload,
  ) {
    return this.userService.registerUser(createUserPayload);
  }

  @Post('login')
  loginUser(@Body(new LoginValidationPipe()) loginPayload: LoginPayload) {
    return this.userService.loginUser(loginPayload);
  }

  @Get()
  @UseGuards(AuthGuard)
  getTodos() {
    return this.userService.getTodos();
  }

  @Post('todo')
  postTodos(@Body() postTodoDTO: PostTodoDTO) {
    return this.userService.postTodos(postTodoDTO);
  }
}
