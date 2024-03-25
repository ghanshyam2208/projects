import * as Joi from 'joi';
import { BadRequestException, PipeTransform } from '@nestjs/common';

export class CreateUserPayload {
  email: string;
  password: string;
}

export const CreateUserPayloadSchema = Joi.object({
  email: Joi.string().required(),
  password: Joi.string().required(),
}).options({
  abortEarly: false,
});

export class CreateUserPayloadValidationPipe implements PipeTransform {
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
