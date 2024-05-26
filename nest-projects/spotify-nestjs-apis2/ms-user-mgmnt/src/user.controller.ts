import {
  Body,
  Controller,
  Get,
  Post,
  Query,
  Request,
  UseGuards,
} from '@nestjs/common';
import {
  CreateUserPayload,
  CreateUserValidationPipe,
  LoginPayload,
  LoginValidationPipe,
} from './user.validation';
import { UserService } from './user.service';
import { PostTodoDTO } from 'proto/todo';
import { AuthGuard, CustomRequest } from './guards/user.guard';

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
  getTodos(@Request() req: CustomRequest) {
    console.log('request object is attached with user', req.user);
    return this.userService.getTodos();
  }

  @Post('todo')
  @UseGuards(AuthGuard)
  postTodos(@Body() postTodoDTO: PostTodoDTO) {
    return this.userService.postTodos(postTodoDTO);
  }

  @Get('verify')
  async verify(
    @Query('userId') userId: string,
    @Query('verificationOtp') verificationOtp: number,
  ) {
    // Access the query parameters
    console.log('userId:', userId);
    console.log('verificationOtp:', verificationOtp);

    // Perform some actions with the parameters
    await this.userService.verifyEmail(userId, verificationOtp);
    // Return a response
    return {
      message: 'Received query parameters',
      userId,
      verificationOtp,
    };
  }
}
