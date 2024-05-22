// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v1.176.0
//   protoc               v3.20.3
// source: proto/auth.proto

/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable } from 'rxjs';

export const protobufPackage = 'auth';

export interface Empty {}

export interface PaginationPayload {
  page: number;
  skip: number;
}

export interface UpdateUserPayload {
  id: string;
  socialMedia: SocialMedia | undefined;
}

export interface FindByIdPayload {
  id: string;
}

export interface Users {
  users: User[];
}

export interface CreateUserPayload {
  username: string;
  password: string;
  age: number;
}

export interface User {
  id: string;
  username: string;
  password: string;
  age: number;
  subscribed: boolean;
  socialMedia: SocialMedia | undefined;
}

export interface SocialMedia {
  twitterUri?: string | undefined;
  fbUri?: string | undefined;
}

export const AUTH_PACKAGE_NAME = 'auth';

export interface UsersServiceClient {
  createUser(request: CreateUserPayload): Observable<User>;

  findAllUsers(request: Empty): Observable<Users>;

  findOneUser(request: FindByIdPayload): Observable<User>;

  updateUser(request: UpdateUserPayload): Observable<User>;

  removeUser(request: FindByIdPayload): Observable<User>;

  queryUsers(request: Observable<PaginationPayload>): Observable<Users>;
}

export interface UsersServiceController {
  createUser(
    request: CreateUserPayload,
  ): Promise<User> | Observable<User> | User;

  findAllUsers(request: Empty): Promise<Users> | Observable<Users> | Users;

  findOneUser(
    request: FindByIdPayload,
  ): Promise<User> | Observable<User> | User;

  updateUser(
    request: UpdateUserPayload,
  ): Promise<User> | Observable<User> | User;

  removeUser(request: FindByIdPayload): Promise<User> | Observable<User> | User;

  queryUsers(request: Observable<PaginationPayload>): Observable<Users>;
}

export function UsersServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = [
      'createUser',
      'findAllUsers',
      'findOneUser',
      'updateUser',
      'removeUser',
    ];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcMethod('UsersService', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
    const grpcStreamMethods: string[] = ['queryUsers'];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcStreamMethod('UsersService', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
  };
}

export const USERS_SERVICE_NAME = 'UsersService';
