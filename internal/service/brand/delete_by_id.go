package brand

import "context"

func (s *brandService) DeleteById(ctx context.Context, id uint32) error {
	err := s.brandRepository.DeleteById(ctx, id)

	return err
}
