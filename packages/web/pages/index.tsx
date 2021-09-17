import styled from '@emotion/styled';

import { PageWrapper } from '../styles/components';

const IndexWrapper = styled(PageWrapper)`
  max-width: var(--content-width-xl);
`;

export default function Desktop() {
  return <IndexWrapper>
    This is the desktop
  </IndexWrapper>;
}