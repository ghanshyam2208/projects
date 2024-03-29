import * as Joi from 'joi';
import { BadRequestException, PipeTransform } from '@nestjs/common';

export class CreateSongPayload {
  title: string;
  artists: string;
  releasedDate: Date;
  duration: Date;
}
// Custom Joi extension for duration validation
const JoiDuration = Joi.extend((joi) => ({
  type: 'duration',
  base: joi.string(),
  messages: {
    'duration.invalid':
      '{{#label}} must be a valid duration string (e.g. "55:59")',
  },
  validate(value, helpers) {
    // Regular expression to match duration strings like "55:59"
    const durationRegex = /^([0-5]?[0-9]):([0-5]?[0-9])$/;
    if (!durationRegex.test(value)) {
      return { value, errors: helpers.error('duration.invalid') };
    }
    return { value };
  },
}));

// Schema for validating duration
const durationSchema = JoiDuration.duration().required();
export const CreateSongPayloadSchema = Joi.object({
  title: Joi.string().required(),
  artists: Joi.array().required(),
  releasedDate: Joi.date().required(),
  duration: durationSchema,
}).options({
  abortEarly: false,
});

export class CreateSongValidationPipe implements PipeTransform {
  public transform(value: CreateSongPayload): CreateSongPayload {
    const result = CreateSongPayloadSchema.validate(value);
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
