import { Controller } from '@nestjs/common';
import { AppService } from './app.service';
import {
  CreateUserPayload,
  FindByIdPayload,
  PaginationPayload,
  UpdateUserPayload,
  Users,
  UsersServiceController,
  UsersServiceControllerMethods,
} from 'proto/auth';
import { randomUUID } from 'crypto';
import { Observable } from 'rxjs';

@Controller()
@UsersServiceControllerMethods()
export class AppController implements UsersServiceController {
  constructor(private readonly appService: AppService) {
    for (let i = 0; i < 100; i++) {
      this.createUser({
        username: randomUUID(),
        age: 1,
        password: randomUUID(),
      });
    }
  }

  createUser(createUserPayload: CreateUserPayload) {
    return this.appService.createUser(createUserPayload);
  }

  findAllUsers() {
    return this.appService.findAllUsers();
  }

  findOneUser(findByIdPayload: FindByIdPayload) {
    return this.appService.findOneUser(findByIdPayload);
  }

  updateUser(updateUserPayload: UpdateUserPayload) {
    return this.appService.updateUser(updateUserPayload);
  }

  removeUser(findByIdPayload: FindByIdPayload) {
    return this.appService.removeUser(findByIdPayload);
  }

  queryUsers(
    paginationPayload: Observable<PaginationPayload>,
  ): Observable<Users> {
    return this.appService.queryUsers(paginationPayload);
  }
}
