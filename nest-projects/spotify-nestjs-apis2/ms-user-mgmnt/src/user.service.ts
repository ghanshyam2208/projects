import {
  BadRequestException,
  Inject,
  Injectable,
  NotFoundException,
  OnModuleInit,
  UnauthorizedException,
} from '@nestjs/common';
import {
  CreateUserPayload,
  LoginPayload,
  sanitizedUserResponse,
} from './user.validation';
import { UserRepository } from './user.repository';
import { CryptoHelper } from './prisma/crypto.helper';
import { PostTodoDTO, TODO_SERVICE_NAME, TodoServiceClient } from 'proto/todo';
import { ClientGrpc } from '@nestjs/microservices';

@Injectable()
export class UserService implements OnModuleInit {
  private todoServiceClient: TodoServiceClient;
  constructor(
    private readonly userRepository: UserRepository,
    private readonly cryptoHelper: CryptoHelper,
    @Inject('todo') private todoGrpcClient: ClientGrpc,
  ) {}

  onModuleInit() {
    this.todoServiceClient =
      this.todoGrpcClient.getService<TodoServiceClient>(TODO_SERVICE_NAME);
  }

  getTodos() {
    return this.todoServiceClient.getTodos({});
  }

  postTodos(postTodoDTO: PostTodoDTO) {
    console.log(postTodoDTO);
    return this.todoServiceClient.postTodo(postTodoDTO);
  }

  async registerUser(createUserPayload: CreateUserPayload) {
    try {
      createUserPayload.password = this.cryptoHelper.encrypt(
        createUserPayload.password,
      );
      const user = await this.userRepository.registerUser(createUserPayload);
      return sanitizedUserResponse(user);
    } catch (error) {
      if (error.code === 'P2002') {
        throw new BadRequestException('User email/username already exists');
      }
      throw error;
    }
  }

  async loginUser(loginPayload: LoginPayload) {
    const user = await this.userRepository.findUserByEmail(loginPayload.email);
    if (!user) {
      throw new NotFoundException(
        `User not found with email ${loginPayload.email}`,
      );
    }
    const decryptedPassword = this.cryptoHelper.decrypt(user.password);
    if (decryptedPassword !== loginPayload.password) {
      throw new UnauthorizedException(`Username or password is wrong`);
    }
    const payload = { sub: user.id, email: user.email };
    return {
      // accessToken: await this.jwtService.signAsync(payload, {
      //   secret: process.env.CRYPTO_SECRET,
      //   expiresIn: '1d',
      // }),
      login: true,
      payload: {
        ...payload,
      },
    };
  }
}
