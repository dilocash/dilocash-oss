/**
 * Copyright (c) 2026 dilocash
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

import React from 'react';
import { centerStyle } from './styles';

import type { VariantProps } from '@gluestack-ui/utils/nativewind-utils';

type ICenterProps = React.ComponentPropsWithoutRef<'div'> &
  VariantProps<typeof centerStyle>;

const Center = React.forwardRef<HTMLDivElement, ICenterProps>(function Center(
  { className, ...props },
  ref
) {
  return (
    <div className={centerStyle({ class: className })} {...props} ref={ref} />
  );
});

Center.displayName = 'Center';

export { Center };
