import { format, parseISO } from "date-fns";

export const getLocalisedTimestamp = (isoTimestampInMicroseconds: string) => {
  return format(
    parseISO(String(isoTimestampInMicroseconds)),
    "dd.MM.yyyy HH:mm",
  );
};
