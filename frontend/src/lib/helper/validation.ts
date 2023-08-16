
// Regular expression pattern to match a valid float with 2 decimals
export function isValidFloat2D(input: string): boolean {
    const pattern = /^\d+(\.\d{1,2})?$/;

    return pattern.test(input);
}