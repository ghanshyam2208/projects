import { BadRequestException, PipeTransform } from '@nestjs/common';
import * as Joi from 'joi';

export class CreateReservationPayload {
  startDate: Date;
  endDate: Date;
  placeId: string;
  invoiceId: string;
}

export const CreateReservationPayloadSchema = Joi.object({
  startDate: Joi.date().required(),
  endDate: Joi.date().required(),
  placeId: Joi.string().required(),
  invoiceId: Joi.string().required(),
}).options({
  abortEarly: false,
});

export class ReservationValidationPipe implements PipeTransform {
  public transform(value: CreateReservationPayload): CreateReservationPayload {
    const result = CreateReservationPayloadSchema.validate(value);
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
