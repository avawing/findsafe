<script module>
  import { defineMeta } from '@storybook/addon-svelte-csf';
  import Button from './Button.svelte';
  import { fn } from '@storybook/test';

  // More on how to set up stories at: https://storybook.js.org/docs/writing-stories
  const { Story } = defineMeta({
    title: 'Example/Button',
    component: Button,
    tags: ['autodocs'],
    argTypes: {
      backgroundColor: { control: 'color' },
      size: {
        control: { type: 'select' },
        options: ['sm', 'md', 'lg', 'xl'],
      },
			color: {
				control: { type: 'select' },
				options: ['primary','tertiary','success','warning'],
			},
    },
    args: {
      onClick: fn(),
    }
  });
</script>

<!-- More on writing stories with args: https://storybook.js.org/docs/writing-stories/args -->
<Story name="Primary" args={{ primary: true, label: 'Button', size: 'md'}} />

<Story name="Secondary" args={{ label: 'Button' }} />

<Story name="Large" args={{ size: 'lg', label: 'Button' }} />

<Story name="Small" args={{ size: 'small', label: 'Button' }} />

<Story name="X-Large" args={{ label: 'Button', size: 'xl' }} />

