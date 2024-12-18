import { withThemeByDataAttribute } from '@storybook/addon-themes';

const preview = {
  parameters: {
    darkMode: { stylePreview: true },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i
      }
    }
  },
  decorators: [
    withThemeByDataAttribute({
      themes: {
        skeleton: 'skeleton',
        crimson: 'crimson',
        modern: 'modern',
        hamlindigo: 'hamlindigo'
      },
      defaultTheme: 'hamlindigo',
      parentSelector: 'body',
      attributeName: 'data-theme'
    })
  ]
};

export default preview;