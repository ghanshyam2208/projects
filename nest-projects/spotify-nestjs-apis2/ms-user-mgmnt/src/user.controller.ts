import { Body, Controller, Get, Post } from '@nestjs/common';
import { CreateUserPayload, CreateUserValidationPipe } from './user.validation';
import { UserService } from './user.service';
import { PostTodoDTO } from 'proto/todo';

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
  loginUser(
    @Body(new CreateUserValidationPipe()) createUserPayload: CreateUserPayload,
  ) {
    return this.userService.loginUser(createUserPayload);
  }

  @Get()
  getTodos() {
    return this.userService.getTodos();
  }

  @Post('todo')
  postTodos(@Body() postTodoDTO: PostTodoDTO) {
    console.log(postTodoDTO);
    return this.userService.postTodos(postTodoDTO);
  }
}
