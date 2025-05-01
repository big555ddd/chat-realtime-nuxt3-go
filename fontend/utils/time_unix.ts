export default function formatUnixToDateTime(unixTimestamp: number): string {
  // Ensure the Unix timestamp is a valid number
  if (isNaN(unixTimestamp) || unixTimestamp <= 0) {
    console.error('Invalid Unix timestamp:', unixTimestamp);
    return 'Invalid Date';
  }

  const date = new Date(unixTimestamp * 1000); // Convert Unix to milliseconds

  // Check if the date is valid
  if (isNaN(date.getTime())) {
    console.error('Invalid date:', date);
    return 'Invalid Date';
  }

  const now = new Date();
  
  // Normalize dates to compare only the date part (ignoring time)
  const currentDate = new Date(now.getFullYear(), now.getMonth(), now.getDate());
  const givenDate = new Date(date.getFullYear(), date.getMonth(), date.getDate());

  const differenceInDays = Math.floor(
    (currentDate.getTime() - givenDate.getTime()) / (1000 * 60 * 60 * 24)
  );

  const timeOptions: Intl.DateTimeFormatOptions = {
    hour: "2-digit",
    minute: "2-digit",
    hour12: true,
  };
  const timeFormatted = new Intl.DateTimeFormat("en-US", timeOptions).format(date);

  if (differenceInDays === 0) {
    // If the date is today, show "HH:MM AM/PM"
    return timeFormatted;
  } else if (differenceInDays === 1) {
    // If the date is yesterday, show "Yesterday HH:MM AM/PM"
    return `Yesterday ${timeFormatted}`;
  } else {
    // For other dates, show "DD MMM YYYY HH:MM AM/PM"
    const dateOptions: Intl.DateTimeFormatOptions = {
      day: "2-digit",
      month: "short",
      year: "numeric",
    };
    const dateFormatted = new Intl.DateTimeFormat("en-US", dateOptions).format(date);
    return `${dateFormatted.toUpperCase()} ${timeFormatted}`;
  }
}
