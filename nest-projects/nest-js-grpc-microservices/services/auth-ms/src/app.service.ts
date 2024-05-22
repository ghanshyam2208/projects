import { Injectable, NotFoundException, OnModuleInit } from '@nestjs/common';
import { randomUUID } from 'crypto';
import {
  CreateUserPayload,
  FindByIdPayload,
  PaginationPayload,
  UpdateUserPayload,
  User,
  Users,
} from 'proto/auth';
import { Observable, Subject } from 'rxjs';

@Injectable()
export class AppService implements OnModuleInit {
  private readonly users: User[] = [];
  onModuleInit() {}

  createUser(createUserPayload: CreateUserPayload) {
    console.log('calling in auth ms user');
    const user: User = {
      ...createUserPayload,
      subscribed: false,
      socialMedia: {},
      id: randomUUID(),
    };
    this.users.push(user);
    console.log('returning user');
    return user;
  }

  findAllUsers() {
    return {
      users: this.users,
    };
  }

  findOneUser(findByIdPayload: FindByIdPayload) {
    return this.users.find((user: User) => {
      return user.id === findByIdPayload.id;
    });
  }

  updateUser(updateUserPayload: UpdateUserPayload) {
    const userIndex = this.users.findIndex(
      (user) => user.id === updateUserPayload.id,
    );

    if (userIndex === -1) {
      this.users[userIndex] = {
        ...this.users[userIndex],
        ...updateUserPayload,
      };
      return this.users[userIndex];
    }

    throw new NotFoundException(
      `user not found with with id ${updateUserPayload.id}`,
    );
  }

  removeUser(findByIdPayload: FindByIdPayload) {
    const userIndex = this.users.findIndex(
      (user) => user.id === findByIdPayload.id,
    );

    if (userIndex === -1) {
      return this.users.splice(userIndex)[0];
    }

    throw new NotFoundException(
      `user not found with with id ${findByIdPayload.id}`,
    );
  }

  queryUsers(
    paginationPayload: Observable<PaginationPayload>,
  ): Observable<Users> {
    const subject = new Subject<Users>();

    const onNext = (paginationPayload: PaginationPayload) => {
      const start = paginationPayload.page * paginationPayload.skip;
      subject.next({
        users: this.users.slice(start, start + paginationPayload.skip),
      });
    };

    const onComplete = () => subject.complete();
    paginationPayload.subscribe({
      next: onNext,
      complete: onComplete,
    });

    return subject.asObservable();
  }
}
