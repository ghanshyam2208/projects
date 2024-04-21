import * as Joi from 'joi';
import { BadRequestException, PipeTransform } from '@nestjs/common';
import { Users } from '@prisma/client';

export class CreateUserPayload {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

export const CreateUserPayloadSchema = Joi.object({
  firstName: Joi.string().required(),
  lastName: Joi.string().required(),
  email: Joi.string().required(),
  password: Joi.string().required(),
}).options({
  abortEarly: false,
});

export class CreateUserValidationPipe implements PipeTransform {
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

// do not add password, or other security related fields
export const UserResponseSchema = Joi.object({
  id: Joi.number(),
  firstName: Joi.string(),
  lastName: Joi.string(),
  email: Joi.string(),
}).unknown(false); // Set unknown(false) to exclude unknown keys

// Validate a user object against the schema
export const sanitizeUserResponse = (user: Users) => {
  const { value: sanitizedUser } = UserResponseSchema.validate(user, {
    stripUnknown: true,
  });

  return sanitizedUser;
};
