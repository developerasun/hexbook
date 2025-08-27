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
        <Button variant="outline" class="w-full">Login with Google</Button>
      </div>
    </form>
  </Card.Content>
</Card.Root>
