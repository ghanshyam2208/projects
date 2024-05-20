export class UserCreatePayload {
  email: string;
  password: string;
}

export class CreateUserEvent {
  constructor(public readonly email: string) {}
}
