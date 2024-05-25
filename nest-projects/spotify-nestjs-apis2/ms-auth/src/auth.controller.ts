import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { PostTodoDTO, TODO_SERVICE_NAME, Todo, Todos } from 'proto/todo';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';
import { AUTH_SERVICE_NAME, AuthToken, GetAuthTokenPayload } from 'proto/auth';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @GrpcMethod(TODO_SERVICE_NAME, 'PostTodo')
  postTodo(postTodoDTO: PostTodoDTO): Promise<Todo> | Observable<Todo> | Todo {
    return this.authService.postTodo(postTodoDTO);
  }

  @GrpcMethod(TODO_SERVICE_NAME, 'GetTodos')
  getTodos(): Promise<Todos> | Observable<Todos> | Todos {
    return this.authService.getTodos();
  }

  @GrpcMethod(AUTH_SERVICE_NAME, 'getAuthToken')
  getAuthToken(
    request: GetAuthTokenPayload,
  ): Promise<AuthToken> | Observable<AuthToken> | AuthToken {
    return this.authService.getAuthToken(request);
  }
}
