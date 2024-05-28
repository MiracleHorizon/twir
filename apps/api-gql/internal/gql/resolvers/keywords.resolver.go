package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/guregu/null"
	model "github.com/satont/twir/libs/gomodels"
	"github.com/twirapp/twir/apps/api-gql/internal/gql/gqlmodel"
)

// KeywordCreate is the resolver for the keywordCreate field.
func (r *mutationResolver) KeywordCreate(ctx context.Context, opts gqlmodel.KeywordCreateInput) (*gqlmodel.Keyword, error) {
	dashboardId, err := r.sessions.GetSelectedDashboard(ctx)
	if err != nil {
		return nil, err
	}

	entity := model.ChannelsKeywords{
		ID:               uuid.NewString(),
		ChannelID:        dashboardId,
		Text:             opts.Text,
		Response:         "",
		Enabled:          true,
		Cooldown:         0,
		CooldownExpireAt: null.Time{},
		IsReply:          true,
		IsRegular:        false,
		Usages:           0,
	}

	if opts.Response.IsSet() && opts.Response.Value() != nil {
		entity.Response = *opts.Response.Value()
	}

	if opts.Enabled.IsSet() {
		entity.Enabled = *opts.Enabled.Value()
	}

	if opts.Cooldown.IsSet() {
		entity.Cooldown = *opts.Cooldown.Value()
	}

	if opts.IsReply.IsSet() {
		entity.IsReply = *opts.IsReply.Value()
	}

	if opts.IsRegularExpression.IsSet() {
		entity.IsRegular = *opts.IsRegularExpression.Value()
	}

	if opts.UsageCount.IsSet() {
		entity.Usages = *opts.UsageCount.Value()
	}

	if err := r.gorm.WithContext(ctx).Create(&entity).Error; err != nil {
		return nil, err
	}

	if err := r.keywordsCacher.Invalidate(ctx, dashboardId); err != nil {
		r.logger.Error("failed to invalidate keywords cache", err)
	}

	return &gqlmodel.Keyword{
		ID:                  entity.ID,
		Text:                entity.Text,
		Response:            &entity.Response,
		Enabled:             entity.Enabled,
		Cooldown:            entity.Cooldown,
		IsReply:             entity.IsReply,
		IsRegularExpression: entity.IsRegular,
		UsageCount:          entity.Usages,
	}, nil
}

// KeywordUpdate is the resolver for the keywordUpdate field.
func (r *mutationResolver) KeywordUpdate(ctx context.Context, id string, opts gqlmodel.KeywordUpdateInput) (*gqlmodel.Keyword, error) {
	dashboardId, err := r.sessions.GetSelectedDashboard(ctx)
	if err != nil {
		return nil, err
	}

	entity := model.ChannelsKeywords{}
	if err := r.gorm.WithContext(ctx).
		Where(`id = ? AND "channelId" = ?`, id, dashboardId).
		First(&entity).Error; err != nil {
		return nil, err
	}

	if opts.Text.IsSet() {
		entity.Text = *opts.Text.Value()
	}

	if opts.Response.IsSet() {
		entity.Response = *opts.Response.Value()
	}

	if opts.Enabled.IsSet() {
		entity.Enabled = *opts.Enabled.Value()
	}

	if opts.Cooldown.IsSet() {
		entity.Cooldown = *opts.Cooldown.Value()
	}

	if opts.IsReply.IsSet() {
		entity.IsReply = *opts.IsReply.Value()
	}

	if opts.IsRegularExpression.IsSet() {
		entity.IsRegular = *opts.IsRegularExpression.Value()
	}

	if opts.UsageCount.IsSet() {
		entity.Usages = *opts.UsageCount.Value()
	}

	if err := r.gorm.WithContext(ctx).Save(&entity).Error; err != nil {
		return nil, err
	}

	if err := r.keywordsCacher.Invalidate(ctx, dashboardId); err != nil {
		r.logger.Error("failed to invalidate keywords cache", err)
	}

	return &gqlmodel.Keyword{
		ID:                  entity.ID,
		Text:                entity.Text,
		Response:            &entity.Response,
		Enabled:             entity.Enabled,
		Cooldown:            entity.Cooldown,
		IsReply:             entity.IsReply,
		IsRegularExpression: entity.IsRegular,
		UsageCount:          entity.Usages,
	}, nil
}

// KeywordRemove is the resolver for the keywordRemove field.
func (r *mutationResolver) KeywordRemove(ctx context.Context, id string) (bool, error) {
	dashboardId, err := r.sessions.GetSelectedDashboard(ctx)
	if err != nil {
		return false, err
	}

	keyword := model.ChannelsKeywords{}
	if err := r.gorm.WithContext(ctx).
		Where(`id = ? AND "channelId" = ?`, id, dashboardId).
		First(&keyword).Error; err != nil {
		return false, fmt.Errorf("keyword not found: %w", err)
	}

	if err := r.gorm.WithContext(ctx).
		Delete(&keyword).Error; err != nil {
		return false, err
	}

	if err := r.keywordsCacher.Invalidate(ctx, dashboardId); err != nil {
		r.logger.Error("failed to invalidate keywords cache", err)
	}

	return true, nil
}

// Keywords is the resolver for the keywords field.
func (r *queryResolver) Keywords(ctx context.Context) ([]gqlmodel.Keyword, error) {
	dashboardId, err := r.sessions.GetSelectedDashboard(ctx)
	if err != nil {
		return nil, err
	}

	var entities []model.ChannelsKeywords
	if err := r.gorm.WithContext(ctx).
		Where(`"channelId" = ?`, dashboardId).
		Order("id ASC").
		Find(&entities).Error; err != nil {
		return nil, err
	}

	var keywords []gqlmodel.Keyword
	for _, entity := range entities {
		keywords = append(
			keywords,
			gqlmodel.Keyword{
				ID:                  entity.ID,
				Text:                entity.Text,
				Response:            &entity.Response,
				Enabled:             entity.Enabled,
				Cooldown:            entity.Cooldown,
				IsReply:             entity.IsReply,
				IsRegularExpression: entity.IsRegular,
				UsageCount:          entity.Usages,
			},
		)
	}

	return keywords, nil
}
