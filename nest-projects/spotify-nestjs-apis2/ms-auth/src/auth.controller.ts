import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { PostTodoDTO, TODO_SERVICE_NAME, Todo, Todos } from 'proto/todo';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';

@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @GrpcMethod(TODO_SERVICE_NAME, 'PostTodo')
  postTodo(postTodoDTO: PostTodoDTO): Promise<Todo> | Observable<Todo> | Todo {
    console.log('received call postTodo');
    return this.authService.postTodo(postTodoDTO);
  }

  @GrpcMethod(TODO_SERVICE_NAME, 'GetTodos')
  getTodos(): Promise<Todos> | Observable<Todos> | Todos {
    console.log('received call getTodos');
    return this.authService.getTodos();
  }
}
