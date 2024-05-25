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
import { AUTH_SERVICE_NAME, AuthServiceClient, AuthToken } from 'proto/auth';

@Injectable()
export class UserService implements OnModuleInit {
  private todoServiceClient: TodoServiceClient;
  private authServiceClient: AuthServiceClient;

  constructor(
    private readonly userRepository: UserRepository,
    private readonly cryptoHelper: CryptoHelper,
    @Inject('MULTI_PACKAGE_LOOKUP_NAME') private todoGrpcClient: ClientGrpc,
  ) {}

  onModuleInit() {
    this.todoServiceClient =
      this.todoGrpcClient.getService<TodoServiceClient>(TODO_SERVICE_NAME);

    this.authServiceClient =
      this.todoGrpcClient.getService<AuthServiceClient>(AUTH_SERVICE_NAME);
  }

  getTodos() {
    return this.todoServiceClient.getTodos({});
  }

  postTodos(postTodoDTO: PostTodoDTO) {
    return this.todoServiceClient.postTodo(postTodoDTO);
  }

  async registerUser(createUserPayload: CreateUserPayload) {
    try {
      createUserPayload.password = this.cryptoHelper.encrypt(
        createUserPayload.password,
      );
      const emailVerificationOtp = (() =>
        Math.floor(Math.random() * 900000) + 100000)().toString();
      const user = await this.userRepository.registerUser({
        ...createUserPayload,
        emailVerificationOtp,
      });
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

    const tokenPayload = (await this.authServiceClient
      .getAuthToken({
        email: user.email,
        userId: user.id,
      })
      .toPromise()) as AuthToken;

    return {
      login: true,
      payload: {
        ...tokenPayload,
      },
    };
  }

  async callAuthVerifyToken(token: string) {
    const verifyTokenResponse = await this.authServiceClient
      .verifyToken({
        accessToken: token,
      })
      .toPromise();

    return verifyTokenResponse;
  }
}
