// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/tlf.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type CryptKeysArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type PublicCanonicalTLFNameAndIDArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type CompleteAndCanonicalizeTlfNameArg struct {
	TlfName string `codec:"tlfName" json:"tlfName"`
}

type TlfInterface interface {
	// CryptKeys returns TLF crypt keys from all generations.
	CryptKeys(context.Context, string) (TLFCryptKeys, error)
	// * tlfCanonicalID returns the canonical name and TLFID for tlfName.
	// * TLFID should not be cached or stored persistently.
	PublicCanonicalTLFNameAndID(context.Context, string) (CanonicalTLFNameAndID, error)
	CompleteAndCanonicalizeTlfName(context.Context, string) (CanonicalTlfName, error)
}

func TlfProtocol(i TlfInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.tlf",
		Methods: map[string]rpc.ServeHandlerDescription{
			"CryptKeys": {
				MakeArg: func() interface{} {
					ret := make([]CryptKeysArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CryptKeysArg)
					if !ok {
						err = rpc.NewTypeError((*[]CryptKeysArg)(nil), args)
						return
					}
					ret, err = i.CryptKeys(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"publicCanonicalTLFNameAndID": {
				MakeArg: func() interface{} {
					ret := make([]PublicCanonicalTLFNameAndIDArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PublicCanonicalTLFNameAndIDArg)
					if !ok {
						err = rpc.NewTypeError((*[]PublicCanonicalTLFNameAndIDArg)(nil), args)
						return
					}
					ret, err = i.PublicCanonicalTLFNameAndID(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"completeAndCanonicalizeTlfName": {
				MakeArg: func() interface{} {
					ret := make([]CompleteAndCanonicalizeTlfNameArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]CompleteAndCanonicalizeTlfNameArg)
					if !ok {
						err = rpc.NewTypeError((*[]CompleteAndCanonicalizeTlfNameArg)(nil), args)
						return
					}
					ret, err = i.CompleteAndCanonicalizeTlfName(ctx, (*typedArgs)[0].TlfName)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TlfClient struct {
	Cli rpc.GenericClient
}

// CryptKeys returns TLF crypt keys from all generations.
func (c TlfClient) CryptKeys(ctx context.Context, tlfName string) (res TLFCryptKeys, err error) {
	__arg := CryptKeysArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.tlf.CryptKeys", []interface{}{__arg}, &res)
	return
}

// * tlfCanonicalID returns the canonical name and TLFID for tlfName.
// * TLFID should not be cached or stored persistently.
func (c TlfClient) PublicCanonicalTLFNameAndID(ctx context.Context, tlfName string) (res CanonicalTLFNameAndID, err error) {
	__arg := PublicCanonicalTLFNameAndIDArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.tlf.publicCanonicalTLFNameAndID", []interface{}{__arg}, &res)
	return
}

func (c TlfClient) CompleteAndCanonicalizeTlfName(ctx context.Context, tlfName string) (res CanonicalTlfName, err error) {
	__arg := CompleteAndCanonicalizeTlfNameArg{TlfName: tlfName}
	err = c.Cli.Call(ctx, "keybase.1.tlf.completeAndCanonicalizeTlfName", []interface{}{__arg}, &res)
	return
}
