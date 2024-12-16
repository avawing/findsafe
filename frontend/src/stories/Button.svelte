<script>
  import '../app.css';
  /**
   * @typedef {Object} Props
   * @property {boolean} [primary] Is this the principal call to action on the page?
   * @property {string} [backgroundColor] What background color to use
   * @property {'sm'|'md'|'lg'|'xl'} [size] How large should the button be?
   * @property {string} label Button contents
   * @property {() => void} [onClick] The onclick event handler
   * @property{'primary'|'tertiary'|'success'|'warning'} [color]
   */

  /** @type {Props} */
  const { primary = false, backgroundColor, size = 'md', label, onClick, color='primary'} = $props();
</script>

<button
  type="button"
  class={['btn variant-ghost', `btn btn-${size} variant-ghost-${color}`].join(' ')}
  style:background-color={backgroundColor}
  onclick={onClick}
>
  {label}
</button>
