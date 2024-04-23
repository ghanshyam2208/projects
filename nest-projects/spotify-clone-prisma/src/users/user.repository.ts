import { Injectable } from '@nestjs/common';
import { PrismaService } from '../common/index.js';
import { CreateUserPayload } from './users.validations';

@Injectable()
export class UsersRepository {
  constructor(private prisma: PrismaService) {}

  createUser(createUserPayload: CreateUserPayload) {
    return this.prisma.users.create({
      data: createUserPayload,
    });
  }

  findUserByEmail(email: string) {
    return this.prisma.users.findFirst({
      where: {
        email,
      },
    });
  }
}
