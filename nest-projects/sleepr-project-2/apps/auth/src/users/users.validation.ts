import { BadRequestException, PipeTransform } from '@nestjs/common';
import * as Joi from 'joi';

export class CreateUserPayload {
  email: string;
  password: string;
}

export const CreateUserPayloadSchema = Joi.object({
  email: Joi.string().required().email(),
  password: Joi.string()
    .required()
    .pattern(
      new RegExp(
        '^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,}$',
      ),
    )
    .message('"password" must be a valid email'),
}).options({
  abortEarly: false,
});

export class CreateUsersValidationPipe implements PipeTransform {
  public transform(value: CreateUserPayload): CreateUserPayload {
    const result = CreateUserPayloadSchema.validate(value);
    if (result.error) {
      const errorMessages = result.error.details
        .map((d) => {
          return d.message.replace(/"/g, "'"); // remove / from string, replace them with '
        })
        .join();
      throw new BadRequestException(errorMessages);
    }
    return value;
  }
}
