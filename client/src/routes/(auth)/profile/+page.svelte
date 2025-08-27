<script>
  import {
    Card,
    CardAction,
    CardDescription,
    CardHeader,
    CardTitle
  } from '$lib/components/ui/card';
  import Spacer from '$lib/components/ui/spacer/spacer.svelte';
  import TypographyH1 from '$lib/components/ui/typography/typographyH1.svelte';
  import * as Avatar from '$lib/components/ui/avatar/index.js';
  import TypographyH2 from '$lib/components/ui/typography/typographyH2.svelte';

  const user = {
    username: 'dummyname',
    email: 'dummay@gmail.com',
    membership: 'free',
    country: 'US',
    cats: [
      {
        name: 'navi',
        age: 2,
        weight: 1.77,
        breed: 'mixed',
        picture: 'none'
      },
      {
        name: 'smith',
        age: 7,
        weight: 4.5,
        breed: 'mixed',
        picture: 'none'
      },
      {
        name: 'brian',
        age: 0.5,
        weight: 1.33,
        breed: 'mixed',
        picture: 'none'
      }
    ]
  };
</script>

<div class="grid grid-cols-1 gap-8 md:grid-cols-[30%_70%]">
  <!-- {/* left */} -->
  <div class="flex max-h-max flex-col justify-between">
    <TypographyH1 text="Me" />
    <Spacer my="my-3" />
    <Card>
      <CardHeader>
        <div class="flex flex-col gap-4">
          <div class="flex items-center justify-center gap-4 border-y-2 py-4">
            <Avatar.Root>
              <Avatar.Image src="https://github.com/shadcn.png" alt="@shadcn" />
              <Avatar.Fallback>CN</Avatar.Fallback>
            </Avatar.Root>
            <span>{user.username}</span>
          </div>
          <div>
            <CardTitle>Eail</CardTitle>
            <CardDescription class="w-full whitespace-normal break-all"
              >{user.email}</CardDescription
            >
          </div>
          <div>
            <CardTitle>Membership</CardTitle>
            <CardDescription>{user.membership}</CardDescription>
          </div>
          <div>
            <CardTitle>Country</CardTitle>
            <CardDescription>{user.country}</CardDescription>
          </div>
        </div>
      </CardHeader>
      <CardAction class="m-auto flex items-center justify-center gap-2">logout here</CardAction>
    </Card>
  </div>

  <!-- right -->
  <div>
    <TypographyH2 text="My Cats" />
    <Spacer my="my-4" />
    {#each user.cats as { name, age, weight, breed, picture }, i (i)}
      <Card>
        <CardHeader>
          <CardTitle>({i + 1}) {name}</CardTitle>
          <CardDescription>
            {picture}
            <p class="my-2 flex w-full justify-between whitespace-normal break-all">
              <span>
                {breed}
                {age}
                {weight}KG
              </span>
            </p>
          </CardDescription>
        </CardHeader>
      </Card>
    {/each}
  </div>
</div>
