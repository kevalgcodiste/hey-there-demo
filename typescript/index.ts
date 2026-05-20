export function greet(name: string): string {
  return `Hello, ${name}!`;
}

if (require.main === module) {
  // eslint-disable-next-line no-console
  console.log(greet("world"));
}
