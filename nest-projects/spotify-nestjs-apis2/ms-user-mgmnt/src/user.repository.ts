import { Injectable } from '@nestjs/common';
import { CreateUserPayload } from './user.validation';
import { PrismaService } from './prisma/prisma.service';

@Injectable()
export class UserRepository {
  constructor(private prisma: PrismaService) {}
  registerUser(createUserPayload: CreateUserPayload) {
    return this.prisma.user.create({
      data: {
        ...createUserPayload,
        isEmailVerified: false,
      },
    });
  }
}
