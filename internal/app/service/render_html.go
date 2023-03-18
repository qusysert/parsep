package service

func (s *Service) RenderHtml(html string) ([]byte, error) {
	return s.render.Render(html)
}
