import { Injectable } from '@nestjs/common';
import { PostTodoDTO, Todo, Todos } from 'proto/todo';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {
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
}
