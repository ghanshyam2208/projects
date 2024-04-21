import { Module } from '@nestjs/common';
import { UsersService } from './users.service';
import { CommonModule } from 'src/common';
import { UsersRepository } from './user.repository';

@Module({
  imports: [CommonModule],
  providers: [UsersService, UsersRepository],
  exports: [UsersService],
})
export class UsersModule {}
