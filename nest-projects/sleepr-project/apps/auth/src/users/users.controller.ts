import { Body, Controller, Post } from '@nestjs/common';
import { UsersDocument } from './users.model';
import { UsersService } from './users.service';
import { CreateUsersDto, UsersValidationPipe } from './users.validations';

@Controller('users')
export class UsersController {
  constructor(private readonly usersService: UsersService) {}

  @Post('')
  createReservation(
    @Body(new UsersValidationPipe())
    createReservationDto: CreateUsersDto,
  ): Promise<UsersDocument> {
    return this.usersService.createUser(createReservationDto);
  }
}
