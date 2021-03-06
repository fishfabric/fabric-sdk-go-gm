/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package gm

import (
	"fmt"
	"github.com/Hyperledger-TWGC/tjfoc-gm/sm2"
	"github.com/tw-bc-group/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp"
	sm2ZH "github.com/tw-bc-group/zhonghuan-ce/sm2"
)

//定义国密SM2 keygen 结构体，实现 KeyGenerator 接口
type gmsm2KeyGenerator struct{}

func (gm *gmsm2KeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	//调用 SM2的注册证书方法
	privKey, err := sm2.GenerateKey(nil)
	if err != nil {
		return nil, fmt.Errorf("failed generating GMSM2 key  [%s]", err)
	}

	return &gmsm2PrivateKey{privKey}, nil
}

//定义中环协同SM2 keygen 结构体，实现 KeyGenerator 接口
type zhcesm2KeyGenerator struct{}

func (gm *zhcesm2KeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	//调用中环协同SM2的注册证书方法
	adapter, err := sm2ZH.CreateSm2KeyAdapter("")
	if err != nil {
		return nil, fmt.Errorf("failed generating ZHCESM2 key  [%s]", err)
	}

	return &zhcesm2KeyAdapter{adapter}, nil
}

//定义国密SM4 keygen 结构体，实现 KeyGenerator 接口
type gmsm4KeyGenerator struct {
	length int
}

func (gm *gmsm4KeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	lowLevelKey, err := GetRandomBytes(gm.length)
	if err != nil {
		return nil, fmt.Errorf("failed generating GMSM4 %d key [%s]", gm.length, err)
	}

	return &gmsm4PrivateKey{lowLevelKey, false}, nil
}
