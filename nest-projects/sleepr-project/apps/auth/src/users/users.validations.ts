import * as Joi from 'joi';
import { BadRequestException, PipeTransform } from '@nestjs/common';

export class CreateUsersDto {
  email: string;
  password: string;
}

export const CreateUsersSchema = Joi.object({
  email: Joi.string().required().email(),
  password: Joi.string().required(),
}).options({
  abortEarly: false,
});

export class UsersValidationPipe implements PipeTransform {
  public transform(value: CreateUsersDto): CreateUsersDto {
    const result = CreateUsersSchema.validate(value);
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
