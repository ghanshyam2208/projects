import { Injectable } from '@nestjs/common';
import { PrismaService } from '../common/index.js';
import { CreateUserPayload } from './users.validations.js';
import { Users } from '@prisma/client';

@Injectable()
export class UsersRepository {
  constructor(private prisma: PrismaService) {}

  createUser(createUserPayload: CreateUserPayload): Promise<Users> {
    return this.prisma.users.create({
      data: createUserPayload,
    });
  }

  findUserByEmail(email: string): Promise<Users> {
    return this.prisma.users.findFirst({
      where: {
        email,
      },
    });
  }
}
