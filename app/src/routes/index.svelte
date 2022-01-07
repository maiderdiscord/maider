<script lang="ts">
  import Layout from '@/components/layout.svelte';
  import {
    Card,
    CardBody,
    CardHeader,
    CardTitle,
    Form,
    FormGroup,
    Input,
    Button,
    Alert,
  } from 'sveltestrap';
  import axios from 'axios';
  import { onMount } from 'svelte';

  let licenseKey = '';
  let error = '';

  const auth = async (licenseKey: string) => {
    const res = await axios.post(
      'https://api.keygen.sh/v1/accounts/54de58e6-e0a0-46dd-8e7a-c4b310c2ebfc/licenses/actions/validate-key',
      {
        meta: {
          key: licenseKey,
          scope: {
            fingerprint: await window.core.getHWID(),
          },
        },
      },
    );
    if (!res.data.meta.valid) {
      error = res.data.meta.detail;
      return;
    }
    localStorage.setItem('licenseKey', licenseKey);
    window.location.href = '/app/';
  };

  onMount(async () => {
    const key = localStorage.getItem('licenseKey');
    await auth(key);
  });
</script>

<Layout>
  <Card color="dark">
    <CardHeader>
      <CardTitle>認証</CardTitle>
      <CardBody>
        <Form>
          <FormGroup>
            <Input placeholder="ライセンスキー" bind:value={licenseKey} />
          </FormGroup>
        </Form>
        <div class="text-center">
          <Button color="primary" on:click={() => auth(licenseKey)}>認証</Button
          >
        </div>
        {#if error.length > 0}
          <Alert class="mt-5" color="danger">{error}</Alert>
        {/if}
      </CardBody>
    </CardHeader>
  </Card>
</Layout>
