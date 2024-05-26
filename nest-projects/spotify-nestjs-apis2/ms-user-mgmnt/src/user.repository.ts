import { Injectable } from '@nestjs/common';
import { CreateUserPayload, UpdateUserPayload } from './user.validation';
import { PrismaService } from './prisma/prisma.service';
// import Prisma, { JsonObject } from '@prisma/client';
import { User } from '@prisma/client';

@Injectable()
export class UserRepository {
  constructor(private prisma: PrismaService) {}
  registerUser(createUserPayload: CreateUserPayload) {
    this.prisma.user.update;
    return this.prisma.user.create({
      data: {
        ...createUserPayload,
        isEmailVerified: false,
        createdAt: new Date(),
      },
    });
  }

  findUserByEmail(email: string): Promise<User> {
    return this.prisma.user.findFirst({
      where: {
        email,
      },
    });
  }

  findUserById(id: string): Promise<User> {
    return this.prisma.user.findFirst({
      where: {
        id,
      },
    });
  }

  updateUser(id: string, data: UpdateUserPayload): Promise<User> {
    return this.prisma.user.update({
      data: {
        ...data,
      },
      where: {
        id,
      },
    });
  }
}
