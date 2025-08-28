<script lang="ts">
  import { Button } from '$lib/components/ui/button/index.js';
  import { Label } from '$lib/components/ui/label/index.js';
  import { Input } from '$lib/components/ui/input/index.js';
  import * as Card from '$lib/components/ui/card/index.js';

  let email = '';
  let password = '';
  function handleSubmit(event: SubmitEvent) {
    event.preventDefault();
    alert(JSON.stringify({ email, password }));
  }
</script>

<Card.Root
  class="absolute left-1/2 top-1/2 w-full max-w-sm -translate-x-1/2 -translate-y-1/2"
  style="transform: translateY(7%);"
>
  <Card.Header>
    <Card.Title>Login to your account</Card.Title>
    <Card.Description>Enter your email below to login to your account</Card.Description>
    <Card.Action>
      <Button variant="link">Sign Up</Button>
    </Card.Action>
  </Card.Header>
  <Card.Content>
    <form onsubmit={handleSubmit}>
      <div class="flex flex-col gap-6">
        <div class="grid gap-2">
          <Label for="email">Email</Label>
          <Input id="email" bind:value={email} type="email" placeholder="m@example.com" required />
        </div>
        <div class="grid gap-2">
          <div class="flex items-center">
            <Label for="password">Password</Label>
            <a href="##" class="ml-auto inline-block text-sm underline-offset-4 hover:underline">
              Forgot your password?
            </a>
          </div>
          <Input id="password" bind:value={password} type="password" required />
        </div>
      </div>
      <div class="mt-6 flex-col gap-2">
        <Button type="submit" class="w-full">Login</Button>
        <Button variant="outline" class="w-full">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path
              d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
              fill="currentColor"
            />
          </svg>
          Login with Google
        </Button>
      </div>
    </form>
  </Card.Content>
</Card.Root>
