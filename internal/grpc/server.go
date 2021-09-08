package grpc

import (
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/api/params"
	"gorm.io/gorm"
	grpc "mutual/pkg/grpc"
)

type grpcServer struct {
	db *gorm.DB
	vk *api.VK
	grpc.UnimplementedMutualServer
}

func NewGrpcServer(db *gorm.DB, vk *api.VK) grpc.MutualServer {
	return &grpcServer{db: db, vk: vk}
}

func (g *grpcServer) GetMutual(_ context.Context, request *grpc.GetMutualRequest) (*grpc.GetMutualResponse, error) {
	targetGroups := g.getGroups(request.Target)
	currentGroups := g.getGroups(request.Current)
	mutualGroups := make([]string, 0)

	for _, group := range targetGroups {
		for _, currentGroup := range currentGroups {
			if currentGroup == group {
				mutualGroups = append(mutualGroups, fmt.Sprint(group))
				continue
			}
		}
	}

	if len(mutualGroups) == 0 {
		return &grpc.GetMutualResponse{Groups: make([]*grpc.Group, 0)}, nil
	}

	groupNames, err := g.vk.GroupsGetByID(params.NewGroupsGetByIDBuilder().GroupIDs(mutualGroups).Params)
	if err != nil {
		return nil, err
	}

	responseGroups := make([]*grpc.Group, 0)
	for _, name := range groupNames {
		responseGroups = append(responseGroups, &grpc.Group{Name: name.Name, Id: int32(name.ID)})
	}

	return &grpc.GetMutualResponse{Groups: responseGroups}, nil
}

func (g *grpcServer) getGroups(id int32) []int {
	//newUser := user.User{VkID: id}
	//err := g.db.First(&newUser).Error
	//if err == nil && len(newUser.Groups) != 0 {
	//	ans := make([]int, 0)
	//	for _, group := range newUser.Groups {
	//		ans = append(ans, int(group))
	//	}
	//
	//	return ans
	//}

	user, err := g.vk.UsersGetSubscriptions(params.NewUsersGetSubscriptionsBuilder().UserID(int(id)).Params)
	if err != nil {
		return make([]int, 0)
	}

	return user.Groups.Items
}
