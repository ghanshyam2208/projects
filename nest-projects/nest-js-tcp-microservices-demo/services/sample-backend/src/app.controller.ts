import { Body, Controller, Get, Post } from '@nestjs/common';
import { AppService } from './app.service';
import { UserCreatePayload } from './user.validation';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getHello(): string {
    return this.appService.getHello();
  }

  @Post()
  createUser(@Body() userCreatePayload: UserCreatePayload) {
    return this.appService.createUser(userCreatePayload);
  }
}
