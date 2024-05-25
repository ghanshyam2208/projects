import { Injectable } from '@nestjs/common';
import {
  AuthToken,
  GetAuthTokenPayload,
  VerifyTokenPayload,
  VerifyTokenResponse,
} from 'proto/auth';
import { PostTodoDTO, Todo, Todos } from 'proto/todo';
import { Observable } from 'rxjs';
import { AuthHelper } from './auth.helper';
import { request } from 'express';

@Injectable()
export class AuthService {
  constructor(private readonly authHelper: AuthHelper) {}
  private readonly todos: Todo[] = [
    {
      description: 'test',
      id: 1,
      isDone: false,
    },
  ];

  postTodo(postTodoDTO: PostTodoDTO): Promise<Todo> | Observable<Todo> | Todo {
    const newTodo: Todo = {
      id: this.todos.length + 1,
      ...postTodoDTO,
    };
    this.todos.push(newTodo);
    return newTodo;
  }

  getTodos(): Promise<Todos> | Observable<Todos> | Todos {
    return {
      Todos: this.todos,
    };
  }

  async getAuthToken(request: GetAuthTokenPayload): Promise<AuthToken> {
    const jwtToken = await this.authHelper.signJwtToken(request);
    return {
      accessToken: jwtToken,
      refreshToken: jwtToken,
    };
  }

  async verifyToken(request: VerifyTokenPayload): Promise<VerifyTokenResponse> {
    const verifyTOkenResponse = await this.authHelper.verifyJwtToken(
      request.accessToken,
    );
    return verifyTOkenResponse;
  }
}
