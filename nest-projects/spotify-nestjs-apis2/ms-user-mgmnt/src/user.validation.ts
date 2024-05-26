import * as Joi from 'joi';
import { BadRequestException, PipeTransform } from '@nestjs/common';
import { User } from '@prisma/client';
const passwordRegex = /^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,40}$/;

export class CreateUserPayload {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
  emailVerificationOtp: string;
}

export class UpdateUserPayload {
  firstName?: string;
  lastName?: string;
  email?: string;
  password?: string;
  emailVerificationOtp?: string;
}

export class LoginPayload {
  email: string;
  password: string;
}

export const CreateUserPayloadSchema = Joi.object({
  firstName: Joi.string().required(),
  lastName: Joi.string().required(),
  email: Joi.string().email().required(),
  password: Joi.string()
    .pattern(new RegExp(passwordRegex))
    .required()
    .messages({
      'string.pattern.base': 'Password must meet the specified requirements',
      'string.empty': 'Password is required',
    }),
}).options({
  abortEarly: false,
});

export const LoginPayloadSchema = Joi.object({
  email: Joi.string().email().required(),
  password: Joi.string()
    .pattern(new RegExp(passwordRegex))
    .required()
    .messages({
      'string.pattern.base': 'Password must meet the specified requirements',
      'string.empty': 'Password is required',
    }),
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

export class LoginValidationPipe implements PipeTransform {
  public transform(value: LoginPayload): LoginPayload {
    const result = LoginPayloadSchema.validate(value);
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
  id: Joi.any(),
  firstName: Joi.any(),
  lastName: Joi.any(),
  email: Joi.any(),
}).unknown(false); // Set unknown(false) to exclude unknown keys

// Validate a user object against the schema
export const sanitizedUserResponse = (user: User) => {
  const { value: sanitizedUser } = UserResponseSchema.validate(user, {
    stripUnknown: true,
  });

  return sanitizedUser;
};
