// Código generado por `go run ./internal/genxsd/generar` a partir de los
// XSD del SAT en internal/genxsd/xsd. NO EDITAR: cambiar el generador.

package cfdi

import "regexp"

// Esquemas de los que salen estas tablas (sha256):
//   - cfd/4/cfdv40.xsd d77f7cc7778c18f5b606809829a479de36d80b7eb3b3aba2111a15084ac03357
//   - cfd/catalogos/catCFDI.xsd 46fe267a4b8f2782a3ae9c6086501a9fcf3c4033c309aaffae3382f95882c537
//   - cfd/tipoDatos/tdCFDI/tdCFDI.xsd 8e3a54e0be41f7eef434904d5e405666be84b9977bc1fe2aa6552e78a0c524e6
//   - cfd/3/cfdv33.xsd 0673ac24d3836a396892104fd5ec08e548556f63544c653993ef8d0f502fb16e
//   - cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd a9c76757be34e6ce75f7a4938f1e59c806589190200bc8ba7ed0e4bbdb504356
//   - cfd/Pagos/Pagos20.xsd 3e9f89786c3248d49b80099acb7f1afd89845a2e8d7f65d3d41f92c890d703be
//   - cfd/catalogos/Pagos/catPagos.xsd d0d821919e5d3287f2dde473b6983405eb8f56710a6f999f20d9398eaf6ddaa7
//   - cfd/nomina/nomina12.xsd bbe69a242879534d58cb1832c9f5e8daeec9f9d725a6f9f1a35f240f9cf13783
//   - cfd/catalogos/Nomina/catNomina.xsd 29466a43c0f92433b947c3eeecfba6ae832af45a379b2b3aee5fee77c064c1ad
//   - cfd/implocal/implocal.xsd 2c443aff32d5122555561b7a715d54c34ef86815ec525605a35ac901b23f851a
//   - cfd/CartaPorte/CartaPorte31.xsd 4a113f7df14c60b00fb2aa81937a8064533a42a31b5eb1531b0d4f178800be5b
//   - cfd/catalogos/ComExt/catComExt.xsd 09dd6366beb1bda8654d3c120101f065b02be93622b54d8f9392d2096d00e38c
//   - cfd/catalogos/CartaPorte/catCartaPorte.xsd aa22699d4bb42a2ecbbf77a5c2965ee622624f8ec596edf5fb372ed6a88ec065
//   - cfd/Pagos/Pagos10.xsd f499ae546ce7b7ab83ed3cebdf221127cb5ec35888f55a96f76aa7631655ccfa
//   - cfd/CartaPorte/CartaPorte30.xsd b3478008a2b83a5c0b7c50fc53f00c9ef228e0a310776ae8e53a89e7c8203802
//   - cfd/CartaPorte/CartaPorte20.xsd 6802aa91c19d0e30257bed7107829d6cd06d9972f603588b267e5232a59914ee
//   - cfd/ComercioExterior20/ComercioExterior20.xsd e0d964043e079ac6bc87ffc24f53401161ffa5684abe8772f7b5596f608521fc
//   - cfd/ComercioExterior11/ComercioExterior11.xsd e91960bb59d3029f17ccf63e1f4591fe927cb0ae432b1d926457267257288774
//   - cfd/hidrocarburospetroliferos.xsd 16e119a039f7baad55f579e369b34d61518ca97dc508dea4038b42ab4791c35c
//   - cfd/catalogos/HidrocarburosPetro/catHidroYPetro.xsd 664a626df5790519b617183b74e2abf436164be6b49e67e6745562eeae04b7e9
//   - cfd/cfdiregistrofiscal/cfdiregistrofiscal.xsd 7e1c1d40e13f15904371d01d1111a99b6c27e3884aecb7113e26ed1698b2cabe
//   - cfd/divisas/divisas.xsd 916979e402ea2fddeff49fdac76b71e416275f4ca083e887c5976ea727a6568c
//   - cfd/EstadoDeCuentaCombustible/ecc12.xsd acaf51905074c8046ddb16149073264c7b193acd270c062451fd925e721e31a4
//   - cfd/catalogos/Combustible/catCombustible.xsd d7960ad41a759be2d9bd5170d9798ed34d27abcf8b91c8841906733e4e981552
//   - cfd/terceros/terceros11.xsd 466585d71c9688dbbfbd3d1fc706f2927af8d792c951c22616a5bf66d624f21d
//   - cfd/ventavehiculos/ventavehiculos11.xsd 7b3da19c901a4f14d8a8b39aa3b86e4eee1eed0da470e94ebbf087664395a6f6
//   - cfd/aerolineas/aerolineas.xsd fcd4f00cd734f8f22f25b3458279eaf058ffb57fd0732a6685fb9aaa25ce0f2b
//   - cfd/leyendasFiscales/leyendasFisc.xsd 7cb31a85b7ae4c745de3ba0cfd7000ce88780cc22c49fbe353628d61565b97e1
//   - cfd/iedu/iedu.xsd fe137c3a317baa59b71d55a966119218afde08d6fe55a8ade922a3cd2c104936
//   - cfd/donat/donat11.xsd 83edbf647aa5cd10f952370fe262113a0fa6d30a642776d9c2a1ca92ffa6692b
//   - cfd/CartaPorte/CartaPorte.xsd a7da39aee945c6c1d568cbcb6acd6b314f3e63177eecf2b6d94b3511a04a1337
//   - cfd/ComercioExterior/ComercioExterior10.xsd 1f509fcdd366094643c1bbdc77b59c0e658a6c30fc150525fd2153b771125143
//   - cfd/EstadoDeCuentaCombustible/ecc11.xsd e6ed4d5210ccf3d3476e1a0e59d848c8684b5db5fd185cb16a2efdd2fced40e3
//   - cfd/nomina/nomina11.xsd f719603f155f9297af50bce5aa90d4d32665f02590b1d2804e94c6badf3b7a0a
//   - cfd/ine/ine11.xsd 4b04cd6c755abccec7271099709bb884eec5c27cf2de1585430617badb3d64ac
//   - cfd/valesdedespensa/valesdedespensa.xsd b3110456ce938e5d54856c3e6c423e01cfb11422d911432a12045248e6bffe50
//   - cfd/consumodecombustibles/consumodecombustibles.xsd 5a2f52c5c0e4b26ccee0972de0886e2cbf8c23239723cefcd1aa85e5b5ea8b2a
//   - cfd/consumodecombustibles/consumodeCombustibles11.xsd a503661a9fda5ff956349d131f2f6baea5de1aefe4a11b7677db29a7f048f3b6
//   - cfd/pagoenespecie/pagoenespecie.xsd 22fdb59903abed6da9c86eca92d512c630a691d619ad1633df1cd7bd43be5c6d
//   - cfd/notariospublicos/notariospublicos.xsd 251d9cf82f8e21a6569016480173d4776a69b35ce509bc0b7ca153cc1d78409e
//   - cfd/servicioparcialconstruccion/servicioparcialconstruccion.xsd bb33b349df059937c2557895b7985136dd9ed8259d64255aad13da27608266ec
//   - cfd/vehiculousado/vehiculousado.xsd 768532eff2037d1d32cc684ca06436a99bc702cbb0d0e083b24264774c3f75ce
//   - cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xsd b22c5a7fe7ba1c0305c6b31f899ad1c8824afb78b10cd4784b2da4978d96ddd9
//   - cfd/acreditamiento/AcreditamientoIEPS10.xsd 7a7a716d3bb5f92fab58df0e67c632dd8a18c6aa7e3f19d0767e034c6bbf780d
//   - cfd/spei/spei.xsd b35a7b35f90d0f16e355ba3f7819d3d9ebecd8ba532480b8850c1271252cb2a9
//   - cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd 577460ca293b1f65d23ab205622d912fdfd0e7bbf97e3d6f8d32c40ba77e346a
//   - cfd/certificadodestruccion/certificadodedestruccion.xsd 27075a12f14f32f40dc47e503dd07b94df2519485f9d591c308ed9cb6f3596b4
//   - cfd/arteantiguedades/obrasarteantiguedades.xsd 114620dea293a0212ffd186381e7d841dd778ff3ad231c87791f65a82ff5e74a
//   - cfd/pfic/pfic.xsd dd754b48731507e2c197d18a711a2ae1c2f145253a048ffe4fd3e15bc1f47dbd
//   - cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd 0b61b3b43c902ddf9e3214c426195f6184bb3b80958ea8e49a50bb50ab4d7812
//   - cfd/catalogos/hidrocarburos/catHidrocarburos.xsd 8279b0949c05724d702686c37fe3e5cdede9dcf8784f210f4f5cdcd1f06b42ef
//   - cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xsd 3b0f45a403309e108553bbb1a92b0cdd49d16b48494c9edaccdfa7056535b243
//   - cfd/ecb/ecb.xsd 07227f751b5a251455901c096434c8f4f6bd31d5f91792e01d9e849be93dafab
//   - cfd/ecc/ecc.xsd 8e372b142b7ef2388d1847a59c5549303e0822abac15fb1ca87dfab96802196a
//   - ../esquemas/retencionpago/2/retencionpagov2.xsd 5cc547890c27d4f8a82bddbfe601c1d214a7914afa8ca65e817eadf2434f3eb7
//   - ../esquemas/retencionpago/1/catalogos/catRetenciones.xsd 111c438270416dfd10745bdc8aed83478ceb98c313d0f72b6be7b9fcd8aac4b3
//   - ../esquemas/retencionpago/1/retencionpagov1.xsd 6cac833a992f6e9b9dc0ad99d08b912819640b66f90f29b96b61eae83453b09d
//   - ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd ae51a8c094256b79c3d272139172a3754e3d5ec92bfa05b56dd42c96625ea391
//   - ../esquemas/retencionpago/1/PlataformasTecnologicas10/catalogos/CatPlataformasTecnologicas.xsd 4f44204abcd76d7ef0edfef6e51323030420443691c7d553240a89f4cbe8a9d4
//   - ../esquemas/retencionpago/1/intereses/intereses.xsd 1b0ad59e3014b9d843194cbd7d6333cd0e68127f8d3f3b2f5d99097a137779ab
//   - ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xsd c121aaaa256d3c41d2f87d34fcccafe7d1b64171ca3672331525f34ddcfecfec
//   - ../esquemas/retencionpago/1/dividendos/dividendos.xsd f2e2934a18ef88654e13df94f78d002cbffe57a145f579e23703ec1eeb8e516d
//   - ../esquemas/retencionpago/1/enajenaciondeacciones/enajenaciondeacciones.xsd b145fab587f5fc404fea666473e707208ca287e4da6f5372273a854e40cab56c
//   - ../esquemas/retencionpago/1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xsd 5352dd611dee1a89c18e5fe1f366f0a6aaebee56db3b0df6260b029e80018740
//   - ../esquemas/retencionpago/1/operacionesconderivados/operacionesconderivados.xsd adaa4a1bc9b6f056c6dce553b512863749e827770ecabdeb9797594da42e49de
//   - ../esquemas/retencionpago/1/intereseshipotecarios/intereseshipotecarios.xsd 08c8f89b8a0ac99b4be37de5c496db59fc62b245c5db67fd25809d64766075df
//   - ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd d36bb0ea058b9623aa84f4287f84e6290f22a3adab1e5f94d9a3cdc299042d25
//   - ../esquemas/retencionpago/1/planesderetiro/planesderetiro.xsd 2d521ca9333c1b2bde88fb2b3c2805d8c9b3de76e4fba0184fde460bc611fcb3
//   - ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xsd 0b53318fb58c9119513b534014583308ff3a4e765b5be779bfe02de686ad14aa
//   - ../esquemas/retencionpago/1/planesderetiro11/catalogos/CatPlanesderetiro.xsd 19f2d1b4bf57ba4f3174d5dea5406602b83b446d3f264a806aa4a37429d33f0c
//   - ../esquemas/retencionpago/1/premios/premios.xsd fa110c0bb5d33d44b3751285ca8b2b96287a514c148a5bf739c3c8886639bf20
//   - ../esquemas/retencionpago/1/sectorfinanciero/sectorfinanciero.xsd f1c27be5a70b02d32a90bfd9c9ad89df992379e14e3342326b3f47857e84e52e
//   - cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd 10eae27389bbb3a1d8f9d58e3cc356a1eff37d99ead226f2780458c94fab0a5e
//   - cfd/donat/donat.xsd 17292955730fe284d8a2b7f2307b21a9eed83c27714ce04791ac924735d6412a
//   - cfd/ventavehiculos/ventavehiculos.xsd 262299c7060593b23a709d73aabc59aa817fefdbce723b713df51e44a5209da2
//   - cfd/ine/ine10.xsd 232a5a571f8a547ccbd9b07d3a32213e8a32e310d8c84d3b6a86fc13f95da364

// globalesXSD son los elementos globales: la raíz del documento y lo que
// puede venir en un xs:any (Complemento, ComplementoConcepto, Addenda).
var globalesXSD = map[nombreXML]*elementoXSD{
	{"http://www.sat.gob.mx/cfd/4", "Comprobante"}:                                                                   eCFDI40_Comprobante,
	{"http://www.sat.gob.mx/cfd/3", "Comprobante"}:                                                                   eCFDI33_Comprobante,
	{"http://www.sat.gob.mx/TimbreFiscalDigital", "TimbreFiscalDigital"}:                                             eTFD11_TimbreFiscalDigital,
	{"http://www.sat.gob.mx/Pagos20", "Pagos"}:                                                                       ePagos20_Pagos,
	{"http://www.sat.gob.mx/nomina12", "Nomina"}:                                                                     eNomina12_Nomina,
	{"http://www.sat.gob.mx/implocal", "ImpuestosLocales"}:                                                           eImpLocal_ImpuestosLocales,
	{"http://www.sat.gob.mx/CartaPorte31", "CartaPorte"}:                                                             eCartaPorte31_CartaPorte,
	{"http://www.sat.gob.mx/Pagos", "Pagos"}:                                                                         ePagos10_Pagos,
	{"http://www.sat.gob.mx/CartaPorte30", "CartaPorte"}:                                                             eCartaPorte30_CartaPorte,
	{"http://www.sat.gob.mx/CartaPorte20", "CartaPorte"}:                                                             eCartaPorte20_CartaPorte,
	{"http://www.sat.gob.mx/ComercioExterior20", "ComercioExterior"}:                                                 eComExt20_ComercioExterior,
	{"http://www.sat.gob.mx/ComercioExterior11", "ComercioExterior"}:                                                 eComExt11_ComercioExterior,
	{"http://www.sat.gob.mx/hidrocarburospetroliferos", "HidroYPetro"}:                                               eHidroYPetro_HidroYPetro,
	{"http://www.sat.gob.mx/registrofiscal", "CFDIRegistroFiscal"}:                                                   eRegistroFiscal_CFDIRegistroFiscal,
	{"http://www.sat.gob.mx/divisas", "Divisas"}:                                                                     eDivisas_Divisas,
	{"http://www.sat.gob.mx/EstadoDeCuentaCombustible12", "EstadoDeCuentaCombustible"}:                               eECC12_EstadoDeCuentaCombustible,
	{"http://www.sat.gob.mx/terceros", "PorCuentadeTerceros"}:                                                        eTerceros11_PorCuentadeTerceros,
	{"http://www.sat.gob.mx/ventavehiculos", "VentaVehiculos"}:                                                       eVentaVehiculos11_VentaVehiculos,
	{"http://www.sat.gob.mx/aerolineas", "Aerolineas"}:                                                               eAerolineas_Aerolineas,
	{"http://www.sat.gob.mx/leyendasFiscales", "LeyendasFiscales"}:                                                   eLeyendasFisc_LeyendasFiscales,
	{"http://www.sat.gob.mx/iedu", "instEducativas"}:                                                                 eIEDU_instEducativas,
	{"http://www.sat.gob.mx/donat", "Donatarias"}:                                                                    eDonat11_Donatarias,
	{"http://www.sat.gob.mx/CartaPorte", "CartaPorte"}:                                                               eCartaPorte10_CartaPorte,
	{"http://www.sat.gob.mx/ComercioExterior", "ComercioExterior"}:                                                   eComExt10_ComercioExterior,
	{"http://www.sat.gob.mx/EstadoDeCuentaCombustible", "EstadoDeCuentaCombustible"}:                                 eECC11_EstadoDeCuentaCombustible,
	{"http://www.sat.gob.mx/nomina", "Nomina"}:                                                                       eNomina11_Nomina,
	{"http://www.sat.gob.mx/ine", "INE"}:                                                                             eINE11_INE,
	{"http://www.sat.gob.mx/valesdedespensa", "ValesDeDespensa"}:                                                     eValesDespensa_ValesDeDespensa,
	{"http://www.sat.gob.mx/consumodecombustibles", "ConsumoDeCombustibles"}:                                         eConsumoComb10_ConsumoDeCombustibles,
	{"http://www.sat.gob.mx/ConsumoDeCombustibles11", "ConsumoDeCombustibles"}:                                       eConsumoComb11_ConsumoDeCombustibles,
	{"http://www.sat.gob.mx/pagoenespecie", "PagoEnEspecie"}:                                                         ePagoEnEspecie_PagoEnEspecie,
	{"http://www.sat.gob.mx/notariospublicos", "NotariosPublicos"}:                                                   eNotarios_NotariosPublicos,
	{"http://www.sat.gob.mx/servicioparcialconstruccion", "parcialesconstruccion"}:                                   eParcialesConstruccion_parcialesconstruccion,
	{"http://www.sat.gob.mx/vehiculousado", "VehiculoUsado"}:                                                         eVehiculoUsado_VehiculoUsado,
	{"http://www.sat.gob.mx/TuristaPasajeroExtranjero", "TuristaPasajeroExtranjero"}:                                 eTuristaPasajero_TuristaPasajeroExtranjero,
	{"http://www.sat.gob.mx/acreditamiento", "acreditamientoIEPS"}:                                                   eAcreditamientoIEPS_acreditamientoIEPS,
	{"http://www.sat.gob.mx/spei", "Complemento_SPEI"}:                                                               eSPEI_ComplementoSPEI,
	{"http://www.sat.gob.mx/renovacionysustitucionvehiculos", "renovacionysustitucionvehiculos"}:                     eRenovacion_renovacionysustitucionvehiculos,
	{"http://www.sat.gob.mx/certificadodestruccion", "certificadodedestruccion"}:                                     eCertDestruccion_certificadodedestruccion,
	{"http://www.sat.gob.mx/arteantiguedades", "obrasarteantiguedades"}:                                              eObrasArte_obrasarteantiguedades,
	{"http://www.sat.gob.mx/pfic", "PFintegranteCoordinado"}:                                                         ePFIC_PFintegranteCoordinado,
	{"http://www.sat.gob.mx/GastosHidrocarburos10", "GastosHidrocarburos"}:                                           eGastosHidro_GastosHidrocarburos,
	{"http://www.sat.gob.mx/IngresosHidrocarburos10", "IngresosHidrocarburos"}:                                       eIngresosHidro_IngresosHidrocarburos,
	{"http://www.sat.gob.mx/ecb", "EstadoDeCuentaBancario"}:                                                          eECB_EstadoDeCuentaBancario,
	{"http://www.sat.gob.mx/ecc", "EstadoDeCuentaCombustible"}:                                                       eECC10_EstadoDeCuentaCombustible,
	{"http://www.sat.gob.mx/esquemas/retencionpago/2", "Retenciones"}:                                                eRet20_Retenciones,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1", "Retenciones"}:                                                eRet10_Retenciones,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", "ServiciosPlataformasTecnologicas"}: eRetPlataformas10_ServiciosPlataformasTecnologicas,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/intereses", "Intereses"}:                                        eRetIntereses_Intereses,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", "Pagosaextranjeros"}:                        eRetPagosExtranjeros_Pagosaextranjeros,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", "Dividendos"}:                                      eRetDividendos_Dividendos,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones", "EnajenaciondeAcciones"}:                eRetEnajenacion_EnajenaciondeAcciones,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso", "Arrendamientoenfideicomiso"}:      eRetArrendamiento_Arrendamientoenfideicomiso,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados", "Operacionesconderivados"}:            eRetDerivados_Operacionesconderivados,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios", "Intereseshipotecarios"}:                eRetHipotecarios_Intereseshipotecarios,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", "Fideicomisonoempresarial"}:          eRetFideicomiso_Fideicomisonoempresarial,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro", "Planesderetiro"}:                              eRetPlanesRetiro10_Planesderetiro,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", "Planesderetiro"}:                            eRetPlanesRetiro11_Planesderetiro,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/premios", "Premios"}:                                            eRetPremios_Premios,
	{"http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero", "SectorFinanciero"}:                          eRetSectorFinanciero_SectorFinanciero,
}

// variantesXSD son elementos globales con el mismo namespace y nombre que
// uno de globalesXSD (ver internal/genxsd.Variantes): se usa la variante
// cuando el elemento trae su atributo con ese valor, y si no, el global.
var variantesXSD = map[nombreXML][]varianteXSD{
	{"http://www.sat.gob.mx/TimbreFiscalDigital", "TimbreFiscalDigital"}: {
		{atributo: "version", valor: "1.0", elem: eTFD10_TimbreFiscalDigital}, // si no, eTFD11_TimbreFiscalDigital
	},
	{"http://www.sat.gob.mx/donat", "Donatarias"}: {
		{atributo: "version", valor: "1.0", elem: eDonat10_Donatarias}, // si no, eDonat11_Donatarias
	},
	{"http://www.sat.gob.mx/ventavehiculos", "VentaVehiculos"}: {
		{atributo: "version", valor: "1.0", elem: eVentaVehiculos10_VentaVehiculos}, // si no, eVentaVehiculos11_VentaVehiculos
	},
	{"http://www.sat.gob.mx/ine", "INE"}: {
		{atributo: "Version", valor: "1.0", elem: eINE10_INE}, // si no, eINE11_INE
	},
}

// Patrones XSD traducidos a RE2 (ver internal/genxsd.traducirPatron).
var (
	patron001 = patronXSD{xsd: `[a-f0-9A-F]{8}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12}`, re: regexp.MustCompile(`^(?:[a-f0-9A-F]{8}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12})$`)}
	patron002 = patronXSD{xsd: `[A-Z&Ñ]{3,4}[0-9]{2}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A]`, re: regexp.MustCompile(`^(?:[A-Z&Ñ]{3,4}[0-9]{2}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A])$`)}
	patron003 = patronXSD{xsd: `[^|]{1,300}`, re: regexp.MustCompile(`^(?:[^|]{1,300})$`)}
	patron004 = patronXSD{xsd: `[0-9]{10}`, re: regexp.MustCompile(`^(?:[0-9]{10})$`)}
	patron005 = patronXSD{xsd: `[0-9]{5}`, re: regexp.MustCompile(`^(?:[0-9]{5})$`)}
	patron006 = patronXSD{xsd: `[0-9]{1,18}(.[0-9]{1,6})?`, re: regexp.MustCompile(`^(?:[0-9]{1,18}(?:[^\n\r][0-9]{1,6})?)$`)}
	patron007 = patronXSD{xsd: `[0-9]{2}  [0-9]{2}  [0-9]{4}  [0-9]{7}`, re: regexp.MustCompile(`^(?:[0-9]{2}  [0-9]{2}  [0-9]{4}  [0-9]{7})$`)}
	patron008 = patronXSD{xsd: `[0-9a-zA-Z]{1,150}`, re: regexp.MustCompile(`^(?:[0-9a-zA-Z]{1,150})$`)}
	patron009 = patronXSD{xsd: `[^|]{1,100}`, re: regexp.MustCompile(`^(?:[^|]{1,100})$`)}
	patron010 = patronXSD{xsd: `[^|]{1,20}`, re: regexp.MustCompile(`^(?:[^|]{1,20})$`)}
	patron011 = patronXSD{xsd: `[^|]{1,1000}`, re: regexp.MustCompile(`^(?:[^|]{1,1000})$`)}
	patron012 = patronXSD{xsd: `[^|]{1,25}`, re: regexp.MustCompile(`^(?:[^|]{1,25})$`)}
	patron013 = patronXSD{xsd: `[^|]{1,40}`, re: regexp.MustCompile(`^(?:[^|]{1,40})$`)}
	patron014 = patronXSD{xsd: `(20[1-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])`, re: regexp.MustCompile(`^(?:(?:20[1-9][0-9])-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01])T(?:(?:[01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]))$`)}
	patron015 = patronXSD{xsd: `[0-9]{20}`, re: regexp.MustCompile(`^(?:[0-9]{20})$`)}
	patron016 = patronXSD{xsd: `[0-9a-zA-Z]{5}`, re: regexp.MustCompile(`^(?:[0-9a-zA-Z]{5})$`)}
	patron017 = patronXSD{xsd: `[^|]{1,254}`, re: regexp.MustCompile(`^(?:[^|]{1,254})$`)}
	patron018 = patronXSD{xsd: `[0-9]{1,150}`, re: regexp.MustCompile(`^(?:[0-9]{1,150})$`)}
	patron019 = patronXSD{xsd: `[A-Z&Ñ]{3}[0-9]{2}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A]`, re: regexp.MustCompile(`^(?:[A-Z&Ñ]{3}[0-9]{2}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A])$`)}
	patron020 = patronXSD{xsd: "([A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|\"|%|&|'|´|-|:|;|>|=|<|@|_|,|\\{|\\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü|Ü){1,150}", re: regexp.MustCompile("^(?:(?:[A-Z]|[a-z]|[0-9]| |Ñ|ñ|!|\"|%|&|'|´|-|:|;|>|=|<|@|_|,|\\{|\\}|`|~|á|é|í|ó|ú|Á|É|Í|Ó|Ú|ü|Ü){1,150})$")}
	patron021 = patronXSD{xsd: `[0-9]{1,18}(.[0-9]{1,2})?`, re: regexp.MustCompile(`^(?:[0-9]{1,18}(?:[^\n\r][0-9]{1,2})?)$`)}
	patron022 = patronXSD{xsd: `([a-f0-9A-F]{8}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12})|([0-9]{3}-[0-9]{2}-[0-9]{9})`, re: regexp.MustCompile(`^(?:(?:[a-f0-9A-F]{8}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12})|(?:[0-9]{3}-[0-9]{2}-[0-9]{9}))$`)}
	patron023 = patronXSD{xsd: `[1-9][0-9]{0,2}`, re: regexp.MustCompile(`^(?:[1-9][0-9]{0,2})$`)}
	patron024 = patronXSD{xsd: `XEXX010101000|[A-Z&Ñ]{3}[0-9]{2}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A]`, re: regexp.MustCompile(`^(?:XEXX010101000|[A-Z&Ñ]{3}[0-9]{2}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A])$`)}
	patron025 = patronXSD{xsd: `[A-Z0-9_]{10,50}`, re: regexp.MustCompile(`^(?:[A-Z0-9_]{10,50})$`)}
	patron026 = patronXSD{xsd: `[A-Z][AEIOUX][A-Z]{2}[0-9]{2}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])[MHX]([ABCMTZ]S|[BCJMOT]C|[CNPST]L|[GNQ]T|[GQS]R|C[MH]|[MY]N|[DH]G|NE|VZ|DF|SP)[BCDFGHJ-NP-TV-Z]{3}[0-9A-Z][0-9]`, re: regexp.MustCompile(`^(?:[A-Z][AEIOUX][A-Z]{2}[0-9]{2}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])[MHX](?:[ABCMTZ]S|[BCJMOT]C|[CNPST]L|[GNQ]T|[GQS]R|C[MH]|[MY]N|[DH]G|NE|VZ|DF|SP)[BCDFGHJ-NP-TV-Z]{3}[0-9A-Z][0-9])$`)}
	patron027 = patronXSD{xsd: `[0-9]{1,3}(.([0-9]{1,3}))?`, re: regexp.MustCompile(`^(?:[0-9]{1,3}(?:[^\n\r](?:[0-9]{1,3}))?)$`)}
	patron028 = patronXSD{xsd: `[0-9]{1,15}`, re: regexp.MustCompile(`^(?:[0-9]{1,15})$`)}
	patron029 = patronXSD{xsd: `((19|20)[0-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])`, re: regexp.MustCompile(`^(?:(?:(?:19|20)[0-9][0-9])-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01]))$`)}
	patron030 = patronXSD{xsd: `P(([1-9][0-9]{0,3})|0)W|P([1-9][0-9]?Y)?(([1-9]|1[012])M)?(0|[1-9]|[12][0-9]|3[01])D`, re: regexp.MustCompile(`^(?:P(?:(?:[1-9][0-9]{0,3})|0)W|P(?:[1-9][0-9]?Y)?(?:(?:[1-9]|1[012])M)?(?:0|[1-9]|[12][0-9]|3[01])D)$`)}
	patron031 = patronXSD{xsd: `[^|]{1,15}`, re: regexp.MustCompile(`^(?:[^|]{1,15})$`)}
	patron032 = patronXSD{xsd: `[0-9]{10,18}`, re: regexp.MustCompile(`^(?:[0-9]{10,18})$`)}
	patron033 = patronXSD{xsd: `[^|]{3,15}`, re: regexp.MustCompile(`^(?:[^|]{3,15})$`)}
	patron034 = patronXSD{xsd: `(([1-9][0-9]{0,4})|[0])(.[0-9]{3})?`, re: regexp.MustCompile(`^(?:(?:(?:[1-9][0-9]{0,4})|[0])(?:[^\n\r][0-9]{3})?)$`)}
	patron035 = patronXSD{xsd: `[^|]{1,55}`, re: regexp.MustCompile(`^(?:[^|]{1,55})$`)}
	patron036 = patronXSD{xsd: `[^|]{1,120}`, re: regexp.MustCompile(`^(?:[^|]{1,120})$`)}
	patron037 = patronXSD{xsd: `[^|]{1,250}`, re: regexp.MustCompile(`^(?:[^|]{1,250})$`)}
	patron038 = patronXSD{xsd: `[^|]{1,30}`, re: regexp.MustCompile(`^(?:[^|]{1,30})$`)}
	patron039 = patronXSD{xsd: `(OR|DE)[0-9]{6}`, re: regexp.MustCompile(`^(?:(?:OR|DE)[0-9]{6})$`)}
	patron040 = patronXSD{xsd: `[^|]{6,40}`, re: regexp.MustCompile(`^(?:[^|]{6,40})$`)}
	patron041 = patronXSD{xsd: `[^|]{1,50}`, re: regexp.MustCompile(`^(?:[^|]{1,50})$`)}
	patron042 = patronXSD{xsd: `[^|]{1,150}`, re: regexp.MustCompile(`^(?:[^|]{1,150})$`)}
	patron043 = patronXSD{xsd: `[^|]{10,30}`, re: regexp.MustCompile(`^(?:[^|]{10,30})$`)}
	patron044 = patronXSD{xsd: `OR[0-9]{6}`, re: regexp.MustCompile(`^(?:OR[0-9]{6})$`)}
	patron045 = patronXSD{xsd: `DE[0-9]{6}`, re: regexp.MustCompile(`^(?:DE[0-9]{6})$`)}
	patron046 = patronXSD{xsd: `[0-9|]{6,7}`, re: regexp.MustCompile(`^(?:[0-9|]{6,7})$`)}
	patron047 = patronXSD{xsd: `([0-9]{1,3}[/]){2}([0-9]{1,3})(cm|plg)`, re: regexp.MustCompile(`^(?:(?:[0-9]{1,3}[/]){2}(?:[0-9]{1,3})(?:cm|plg))$`)}
	patron048 = patronXSD{xsd: `[^|]{1,240}`, re: regexp.MustCompile(`^(?:[^|]{1,240})$`)}
	patron049 = patronXSD{xsd: `[^|]{1,10}`, re: regexp.MustCompile(`^(?:[^|]{1,10})$`)}
	patron050 = patronXSD{xsd: `[^|]{1,6}`, re: regexp.MustCompile(`^(?:[^|]{1,6})$`)}
	patron051 = patronXSD{xsd: `[^|]{1,80}`, re: regexp.MustCompile(`^(?:[^|]{1,80})$`)}
	patron052 = patronXSD{xsd: `[^|]{1,60}`, re: regexp.MustCompile(`^(?:[^|]{1,60})$`)}
	patron053 = patronXSD{xsd: `[^|]{1,600}`, re: regexp.MustCompile(`^(?:[^|]{1,600})$`)}
	patron054 = patronXSD{xsd: `[^(?!.*\s)-]{5,7}`, re: regexp.MustCompile(`^(?:[^(?!.*\x20\t\n\r)\-]{5,7})$`)}
	patron055 = patronXSD{xsd: `(19[0-9]{2}|20[0-9]{2})`, re: regexp.MustCompile(`^(?:(?:19[0-9]{2}|20[0-9]{2}))$`)}
	patron056 = patronXSD{xsd: `[^|]{3,50}`, re: regexp.MustCompile(`^(?:[^|]{3,50})$`)}
	patron057 = patronXSD{xsd: `[^|]{3,30}`, re: regexp.MustCompile(`^(?:[^|]{3,30})$`)}
	patron058 = patronXSD{xsd: `[^|]{11,15}`, re: regexp.MustCompile(`^(?:[^|]{11,15})$`)}
	patron059 = patronXSD{xsd: `[^|]{5,20}`, re: regexp.MustCompile(`^(?:[^|]{5,20})$`)}
	patron060 = patronXSD{xsd: `[C]{3}[a-f0-9A-F]{5}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12}`, re: regexp.MustCompile(`^(?:[C]{3}[a-f0-9A-F]{5}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12})$`)}
	patron061 = patronXSD{xsd: `[^|]{7,30}`, re: regexp.MustCompile(`^(?:[^|]{7,30})$`)}
	patron062 = patronXSD{xsd: `IMO[0-9]{7}`, re: regexp.MustCompile(`^(?:IMO[0-9]{7})$`)}
	patron063 = patronXSD{xsd: `[^|]{1,19}`, re: regexp.MustCompile(`^(?:[^|]{1,19})$`)}
	patron064 = patronXSD{xsd: `([A-Z]|[0-9]|-){5,11}`, re: regexp.MustCompile(`^(?:(?:[A-Z]|[0-9]|-){5,11})$`)}
	patron065 = patronXSD{xsd: `[^|]{12,15}`, re: regexp.MustCompile(`^(?:[^|]{12,15})$`)}
	patron066 = patronXSD{xsd: `[^|]{6,15}`, re: regexp.MustCompile(`^(?:[^|]{6,15})$`)}
	patron067 = patronXSD{xsd: `[^|]{6,16}`, re: regexp.MustCompile(`^(?:[^|]{6,16})$`)}
	patron068 = patronXSD{xsd: `[^|]{3,20}`, re: regexp.MustCompile(`^(?:[^|]{3,20})$`)}
	patron069 = patronXSD{xsd: `([A-Z]|[0-9]|-|){5,10}`, re: regexp.MustCompile(`^(?:(?:[A-Z]|[0-9]|-|){5,10})$`)}
	patron070 = patronXSD{xsd: `[^|]{1,35}`, re: regexp.MustCompile(`^(?:[^|]{1,35})$`)}
	patron071 = patronXSD{xsd: `[0-9]{1,14}(.([0-9]{1,3}))?`, re: regexp.MustCompile(`^(?:[0-9]{1,14}(?:[^\n\r](?:[0-9]{1,3}))?)$`)}
	patron072 = patronXSD{xsd: `[0-9]{1,16}(.([0-9]{1,6}))?`, re: regexp.MustCompile(`^(?:[0-9]{1,16}(?:[^\n\r](?:[0-9]{1,6}))?)$`)}
	patron073 = patronXSD{xsd: `[0-9]{1,16}(.([0-9]{1,4}))?`, re: regexp.MustCompile(`^(?:[0-9]{1,16}(?:[^\n\r](?:[0-9]{1,4}))?)$`)}
	patron074 = patronXSD{xsd: `[0-9]{1,16}(.([0-9]{1,2}))?`, re: regexp.MustCompile(`^(?:[0-9]{1,16}(?:[^\n\r](?:[0-9]{1,2}))?)$`)}
	patron075 = patronXSD{xsd: `[0-9]{16}`, re: regexp.MustCompile(`^(?:[0-9]{16})$`)}
	patron076 = patronXSD{xsd: `[0-9]{1,14}(.([0-9]{1,2}))`, re: regexp.MustCompile(`^(?:[0-9]{1,14}(?:[^\n\r](?:[0-9]{1,2})))$`)}
	patron077 = patronXSD{xsd: `[0-9]{1,14}(.([0-9]{1,3}))`, re: regexp.MustCompile(`^(?:[0-9]{1,14}(?:[^\n\r](?:[0-9]{1,3})))$`)}
	patron078 = patronXSD{xsd: `[A-Z][A,E,I,O,U,X][A-Z]{2}[0-9]{2}[0-1][0-9][0-3][0-9][M,H,X][A-Z]{2}[B,C,D,F,G,H,J,K,L,M,N,Ñ,P,Q,R,S,T,V,W,X,Y,Z]{3}[0-9,A-Z][0-9]`, re: regexp.MustCompile(`^(?:[A-Z][A,E,I,O,U,X][A-Z]{2}[0-9]{2}[0-1][0-9][0-3][0-9][M,H,X][A-Z]{2}[B,C,D,F,G,H,J,K,L,M,N,Ñ,P,Q,R,S,T,V,W,X,Y,Z]{3}[0-9,A-Z][0-9])$`)}
	patron079 = patronXSD{xsd: `[A-Z,Ñ,&]{3,4}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z,0-9]?[A-Z,0-9]?[0-9,A-Z]?`, re: regexp.MustCompile(`^(?:[A-Z,Ñ,&]{3,4}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z,0-9]?[A-Z,0-9]?[0-9,A-Z]?)$`)}
	patron080 = patronXSD{xsd: `[0-9]{2}[/]{1}[0-9]{2}[/]{1}[0-9]{2}cm|[0-9]{2}[/]{1}[0-9]{2}[/]{1}[0-9]{2}plg`, re: regexp.MustCompile(`^(?:[0-9]{2}[/]{1}[0-9]{2}[/]{1}[0-9]{2}cm|[0-9]{2}[/]{1}[0-9]{2}[/]{1}[0-9]{2}plg)$`)}
	patron081 = patronXSD{xsd: `[^(?!.*\s)-]{6,7}`, re: regexp.MustCompile(`^(?:[^(?!.*\x20\t\n\r)\-]{6,7})$`)}
	patron082 = patronXSD{xsd: `[A-Z&Ñ]{4}[0-9]{2}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A]`, re: regexp.MustCompile(`^(?:[A-Z&Ñ]{4}[0-9]{2}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])[A-Z0-9]{2}[0-9A])$`)}
	patron083 = patronXSD{xsd: `[A-Z][AEIOUX][A-Z]{2}[0-9]{2}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])[MH]([ABCMTZ]S|[BCJMOT]C|[CNPST]L|[GNQ]T|[GQS]R|C[MH]|[MY]N|[DH]G|NE|VZ|DF|SP)[BCDFGHJ-NP-TV-Z]{3}[0-9A-Z][0-9]`, re: regexp.MustCompile(`^(?:[A-Z][AEIOUX][A-Z]{2}[0-9]{2}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])[MH](?:[ABCMTZ]S|[BCJMOT]C|[CNPST]L|[GNQ]T|[GQS]R|C[MH]|[MY]N|[DH]G|NE|VZ|DF|SP)[BCDFGHJ-NP-TV-Z]{3}[0-9A-Z][0-9])$`)}
	patron084 = patronXSD{xsd: `[a-f0-9A-F]{8}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12}|[A-Za-z0-9,#,+,%,(&) ]{6,40}`, re: regexp.MustCompile(`^(?:[a-f0-9A-F]{8}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{4}-[a-f0-9A-F]{12}|[A-Za-z0-9,#,+,%,(&) ]{6,40})$`)}
	patron085 = patronXSD{xsd: `([0-9]{4})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])`, re: regexp.MustCompile(`^(?:(?:[0-9]{4})-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01])T(?:(?:[01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9]))$`)}
	patron086 = patronXSD{xsd: `[0-9]{3}`, re: regexp.MustCompile(`^(?:[0-9]{3})$`)}
	patron087 = patronXSD{xsd: `[A-Z][A,E,I,O,U,X][A-Z]{2}[0-9]{2}[0-1][0-9][0-3][0-9][M,H][A-Z]{2}[B,C,D,F,G,H,J,K,L,M,N,Ñ,P,Q,R,S,T,V,W,X,Y,Z]{3}[0-9,A-Z][0-9]`, re: regexp.MustCompile(`^(?:[A-Z][A,E,I,O,U,X][A-Z]{2}[0-9]{2}[0-1][0-9][0-3][0-9][M,H][A-Z]{2}[B,C,D,F,G,H,J,K,L,M,N,Ñ,P,Q,R,S,T,V,W,X,Y,Z]{3}[0-9,A-Z][0-9])$`)}
	patron088 = patronXSD{xsd: `[0-9]{18}`, re: regexp.MustCompile(`^(?:[0-9]{18})$`)}
	patron089 = patronXSD{xsd: `[0-9]{1,14}(.([0-9]{3}))`, re: regexp.MustCompile(`^(?:[0-9]{1,14}(?:[^\n\r](?:[0-9]{3})))$`)}
	patron090 = patronXSD{xsd: `[A-ZÑ&]{3}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z0-9]?[A-Z0-9]?[0-9A-Z]-(18|19|20)\d\d(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])-[0-9]{3}`, re: regexp.MustCompile(`^(?:[A-ZÑ&]{3}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z0-9]?[A-Z0-9]?[0-9A-Z]-(?:18|19|20)\p{Nd}\p{Nd}(?:0[1-9]|1[012])(?:0[1-9]|[12][0-9]|3[01])-[0-9]{3})$`)}
	patron091 = patronXSD{xsd: `PE-[0-9]{2}-[0-9]{5}`, re: regexp.MustCompile(`^(?:PE-[0-9]{2}-[0-9]{5})$`)}
	patron092 = patronXSD{xsd: `[0-9]{4}`, re: regexp.MustCompile(`^(?:[0-9]{4})$`)}
	patron093 = patronXSD{xsd: `[0-9]{2}`, re: regexp.MustCompile(`^(?:[0-9]{2})$`)}
	patron094 = patronXSD{xsd: `[A-Z,Ñ]{3}`, re: regexp.MustCompile(`^(?:[A-Z,Ñ]{3})$`)}
	patron095 = patronXSD{xsd: `((000[1-9])|(00[1-9][0-9])|(0[1-9][0-9]{2})|([1-9][0-9]{3}))-((0[1-9])|(1[012]))-((0[1-9])|([12][0-9])|(3[01]))`, re: regexp.MustCompile(`^(?:(?:(?:000[1-9])|(?:00[1-9][0-9])|(?:0[1-9][0-9]{2})|(?:[1-9][0-9]{3}))-(?:(?:0[1-9])|(?:1[012]))-(?:(?:0[1-9])|(?:[12][0-9])|(?:3[01])))$`)}
	patron096 = patronXSD{xsd: `(([01][0-9])|(2[0-3]))(:[0-5][0-9]){2}(\.[0-9]+)?`, re: regexp.MustCompile(`^(?:(?:(?:[01][0-9])|(?:2[0-3]))(?::[0-5][0-9]){2}(?:\.[0-9]+)?)$`)}
	patron097 = patronXSD{xsd: `[A-Z,Ñ,&]{3,4}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z,0-9][A-Z,0-9][0-9,A-Z]`, re: regexp.MustCompile(`^(?:[A-Z,Ñ,&]{3,4}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z,0-9][A-Z,0-9][0-9,A-Z])$`)}
	patron098 = patronXSD{xsd: `([0-9]{4})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])`, re: regexp.MustCompile(`^(?:(?:[0-9]{4})-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01]))$`)}
	patron099 = patronXSD{xsd: `[A-Z,Ñ,&]{3,4}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z,0-9]{2}[0-9,A]`, re: regexp.MustCompile(`^(?:[A-Z,Ñ,&]{3,4}[0-9]{2}[0-1][0-9][0-3][0-9][A-Z,0-9]{2}[0-9,A])$`)}
	patron100 = patronXSD{xsd: `[0-9a-zA-Z]{1,20}`, re: regexp.MustCompile(`^(?:[0-9a-zA-Z]{1,20})$`)}
	patron101 = patronXSD{xsd: `-?([1-9][0-9]{3,}|0[0-9]{3})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](\+|-)((0[0-9]|1[0-3]):[0-5][0-9]|14:00))`, re: regexp.MustCompile(`^(?:-?(?:[1-9][0-9]{3,}|0[0-9]{3})-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12][0-9]|3[01])T(?:(?:[01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9](?:\+|-)(?:(?:0[0-9]|1[0-3]):[0-5][0-9]|14:00)))$`)}
	patron102 = patronXSD{xsd: `\d{1,3}\.\d{1,6}`, re: regexp.MustCompile(`^(?:\p{Nd}{1,3}\.\p{Nd}{1,6})$`)}
)

var xsstring = &tipoXSD{nombre: "xs:string", prim: "string", ws: "preserve"}
var tCatCFDI_cPeriodicidad = &tipoXSD{nombre: "c_Periodicidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cPeriodicidad},
}}
var tCatCFDI_cMeses = &tipoXSD{nombre: "c_Meses", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cMeses},
}}
var xsshort = &tipoXSD{nombre: "xs:short", prim: "short", ws: "collapse"}
var tCFDI40_Comprobante_InformacionGlobal_Anio = &tipoXSD{nombre: "CFDI40 Comprobante/InformacionGlobal/@Año", prim: "short", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "2019"},
}}
var tCFDI40_Comprobante_CfdiRelacionados_CfdiRelacionado_UUID = &tipoXSD{nombre: "CFDI40 Comprobante/CfdiRelacionados/CfdiRelacionado/@UUID", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tCatCFDI_cTipoRelacion = &tipoXSD{nombre: "c_TipoRelacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cTipoRelacion},
}}
var tTdCFDI_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tCFDI40_Comprobante_Emisor_Nombre = &tipoXSD{nombre: "CFDI40 Comprobante/Emisor/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tCatCFDI_cRegimenFiscal = &tipoXSD{nombre: "c_RegimenFiscal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cRegimenFiscal},
}}
var tCFDI40_Comprobante_Emisor_FacAtrAdquirente = &tipoXSD{nombre: "CFDI40 Comprobante/Emisor/@FacAtrAdquirente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 10},
	{tipo: "pattern", patrones: []patronXSD{patron004}},
}}
var tCFDI40_Comprobante_Receptor_Nombre = &tipoXSD{nombre: "CFDI40 Comprobante/Receptor/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tCFDI40_Comprobante_Receptor_DomicilioFiscalReceptor = &tipoXSD{nombre: "CFDI40 Comprobante/Receptor/@DomicilioFiscalReceptor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 5},
	{tipo: "pattern", patrones: []patronXSD{patron005}},
}}
var tCatCFDI_cPais = &tipoXSD{nombre: "c_Pais", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cPais},
}}
var tCFDI40_Comprobante_Receptor_NumRegIdTrib = &tipoXSD{nombre: "CFDI40 Comprobante/Receptor/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tCatCFDI_cUsoCFDI = &tipoXSD{nombre: "c_UsoCFDI", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cUsoCFDI},
}}
var xsdecimal = &tipoXSD{nombre: "xs:decimal", prim: "decimal", ws: "collapse"}
var tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_Base = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Impuestos/Traslados/Traslado/@Base", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCatCFDI_cImpuesto = &tipoXSD{nombre: "c_Impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cImpuesto},
}}
var tCatCFDI_cTipoFactor = &tipoXSD{nombre: "c_TipoFactor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cTipoFactor},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_TasaOCuota = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Impuestos/Traslados/Traslado/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000000"},
}}
var tTdCFDI_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "pattern", patrones: []patronXSD{patron006}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_Base = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Impuestos/Retenciones/Retencion/@Base", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_TasaOCuota = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Impuestos/Retenciones/Retencion/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "fractionDigits", n: 6},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_ACuentaTerceros_NombreACuentaTerceros = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/ACuentaTerceros/@NombreACuentaTerceros", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_ACuentaTerceros_DomicilioFiscalACuentaTerceros = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/ACuentaTerceros/@DomicilioFiscalACuentaTerceros", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 5},
	{tipo: "pattern", patrones: []patronXSD{patron005}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_InformacionAduanera_NumeroPedimento = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/InformacionAduanera/@NumeroPedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_CuentaPredial_Numero = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/CuentaPredial/@Numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron008}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera_NumeroPedimento = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Parte/InformacionAduanera/@NumeroPedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCatCFDI_cClaveProdServ = &tipoXSD{nombre: "c_ClaveProdServ", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cClaveProdServ},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Parte_NoIdentificacion = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Parte/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Parte_Cantidad = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Parte/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Parte_Unidad = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Parte/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Parte_Descripcion = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/Parte/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_NoIdentificacion = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Cantidad = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCatCFDI_cClaveUnidad = &tipoXSD{nombre: "c_ClaveUnidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cClaveUnidad},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Unidad = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCFDI40_Comprobante_Conceptos_Concepto_Descripcion = &tipoXSD{nombre: "CFDI40 Comprobante/Conceptos/Concepto/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCatCFDI_cObjetoImp = &tipoXSD{nombre: "c_ObjetoImp", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cObjetoImp},
}}
var tCFDI40_Comprobante_Impuestos_Traslados_Traslado_TasaOCuota = &tipoXSD{nombre: "CFDI40 Comprobante/Impuestos/Traslados/Traslado/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "fractionDigits", n: 6},
}}
var tCFDI40_Comprobante_Version = &tipoXSD{nombre: "CFDI40 Comprobante/@Version", prim: "string", ws: "collapse"}
var tCFDI40_Comprobante_Serie = &tipoXSD{nombre: "CFDI40 Comprobante/@Serie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron012}},
}}
var tCFDI40_Comprobante_Folio = &tipoXSD{nombre: "CFDI40 Comprobante/@Folio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron013}},
}}
var xsdateTime = &tipoXSD{nombre: "xs:dateTime", prim: "dateTime", ws: "collapse"}
var tTdCFDI_tFechaH = &tipoXSD{nombre: "t_FechaH", prim: "dateTime", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron014}},
}}
var tCFDI40_Comprobante_Sello = &tipoXSD{nombre: "CFDI40 Comprobante/@Sello", prim: "string", ws: "collapse"}
var tCatCFDI_cFormaPago = &tipoXSD{nombre: "c_FormaPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cFormaPago},
}}
var tCFDI40_Comprobante_NoCertificado = &tipoXSD{nombre: "CFDI40 Comprobante/@NoCertificado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron015}},
}}
var tCFDI40_Comprobante_Certificado = &tipoXSD{nombre: "CFDI40 Comprobante/@Certificado", prim: "string", ws: "collapse"}
var tCFDI40_Comprobante_CondicionesDePago = &tipoXSD{nombre: "CFDI40 Comprobante/@CondicionesDePago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCatCFDI_cMoneda = &tipoXSD{nombre: "c_Moneda", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cMoneda},
}}
var tCFDI40_Comprobante_TipoCambio = &tipoXSD{nombre: "CFDI40 Comprobante/@TipoCambio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCatCFDI_cTipoDeComprobante = &tipoXSD{nombre: "c_TipoDeComprobante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cTipoDeComprobante},
}}
var tCatCFDI_cExportacion = &tipoXSD{nombre: "c_Exportacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cExportacion},
}}
var tCatCFDI_cMetodoPago = &tipoXSD{nombre: "c_MetodoPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cMetodoPago},
}}
var tCatCFDI_cCodigoPostal = &tipoXSD{nombre: "c_CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cCodigoPostal},
}}
var tCFDI40_Comprobante_Confirmacion = &tipoXSD{nombre: "CFDI40 Comprobante/@Confirmacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 5},
	{tipo: "pattern", patrones: []patronXSD{patron016}},
}}
var tCFDI33_Comprobante_CfdiRelacionados_CfdiRelacionado_UUID = &tipoXSD{nombre: "CFDI33 Comprobante/CfdiRelacionados/CfdiRelacionado/@UUID", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tCFDI33_Comprobante_Emisor_Nombre = &tipoXSD{nombre: "CFDI33 Comprobante/Emisor/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCFDI33_Comprobante_Receptor_Nombre = &tipoXSD{nombre: "CFDI33 Comprobante/Receptor/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCFDI33_Comprobante_Receptor_NumRegIdTrib = &tipoXSD{nombre: "CFDI33 Comprobante/Receptor/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_Base = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Impuestos/Traslados/Traslado/@Base", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_TasaOCuota = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Impuestos/Traslados/Traslado/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000000"},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_Base = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Impuestos/Retenciones/Retencion/@Base", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_TasaOCuota = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Impuestos/Retenciones/Retencion/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "fractionDigits", n: 6},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_InformacionAduanera_NumeroPedimento = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/InformacionAduanera/@NumeroPedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_CuentaPredial_Numero = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/CuentaPredial/@Numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron018}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera_NumeroPedimento = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Parte/InformacionAduanera/@NumeroPedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Parte_NoIdentificacion = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Parte/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Parte_Cantidad = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Parte/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Parte_Unidad = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Parte/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Parte_Descripcion = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/Parte/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_NoIdentificacion = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Cantidad = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Unidad = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCFDI33_Comprobante_Conceptos_Concepto_Descripcion = &tipoXSD{nombre: "CFDI33 Comprobante/Conceptos/Concepto/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCFDI33_Comprobante_Impuestos_Traslados_Traslado_TasaOCuota = &tipoXSD{nombre: "CFDI33 Comprobante/Impuestos/Traslados/Traslado/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "fractionDigits", n: 6},
}}
var tCFDI33_Comprobante_Version = &tipoXSD{nombre: "CFDI33 Comprobante/@Version", prim: "string", ws: "collapse"}
var tCFDI33_Comprobante_Serie = &tipoXSD{nombre: "CFDI33 Comprobante/@Serie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron012}},
}}
var tCFDI33_Comprobante_Folio = &tipoXSD{nombre: "CFDI33 Comprobante/@Folio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron013}},
}}
var tCFDI33_Comprobante_Sello = &tipoXSD{nombre: "CFDI33 Comprobante/@Sello", prim: "string", ws: "collapse"}
var tCFDI33_Comprobante_NoCertificado = &tipoXSD{nombre: "CFDI33 Comprobante/@NoCertificado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron015}},
}}
var tCFDI33_Comprobante_Certificado = &tipoXSD{nombre: "CFDI33 Comprobante/@Certificado", prim: "string", ws: "collapse"}
var tCFDI33_Comprobante_CondicionesDePago = &tipoXSD{nombre: "CFDI33 Comprobante/@CondicionesDePago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCFDI33_Comprobante_TipoCambio = &tipoXSD{nombre: "CFDI33 Comprobante/@TipoCambio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCFDI33_Comprobante_Confirmacion = &tipoXSD{nombre: "CFDI33 Comprobante/@Confirmacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 5},
	{tipo: "pattern", patrones: []patronXSD{patron016}},
}}
var xsanySimpleType = &tipoXSD{nombre: "xs:anySimpleType", prim: "anySimpleType", ws: "preserve"}
var tTFD11_TimbreFiscalDigital_UUID = &tipoXSD{nombre: "TFD11 TimbreFiscalDigital/@UUID", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tTdCFDI_tRFCPM = &tipoXSD{nombre: "t_RFC_PM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "pattern", patrones: []patronXSD{patron019}},
}}
var tTFD11_TimbreFiscalDigital_Leyenda = &tipoXSD{nombre: "TFD11 TimbreFiscalDigital/@Leyenda", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron020}},
}}
var tTFD11_TimbreFiscalDigital_SelloCFD = &tipoXSD{nombre: "TFD11 TimbreFiscalDigital/@SelloCFD", prim: "string", ws: "collapse"}
var tTFD11_TimbreFiscalDigital_NoCertificadoSAT = &tipoXSD{nombre: "TFD11 TimbreFiscalDigital/@NoCertificadoSAT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron015}},
}}
var tTFD11_TimbreFiscalDigital_SelloSAT = &tipoXSD{nombre: "TFD11 TimbreFiscalDigital/@SelloSAT", prim: "string", ws: "collapse"}
var tTdCFDI_tImporteMXN = &tipoXSD{nombre: "t_ImporteMXN", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "pattern", patrones: []patronXSD{patron021}},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR_RetencionDR_BaseDR = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/ImpuestosDR/RetencionesDR/RetencionDR/@BaseDR", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000001"},
	{tipo: "fractionDigits", n: 6},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR_RetencionDR_TasaOCuotaDR = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/ImpuestosDR/RetencionesDR/RetencionDR/@TasaOCuotaDR", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "fractionDigits", n: 6},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR_TrasladoDR_BaseDR = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/ImpuestosDR/TrasladosDR/TrasladoDR/@BaseDR", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000001"},
	{tipo: "fractionDigits", n: 6},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR_TrasladoDR_TasaOCuotaDR = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/ImpuestosDR/TrasladosDR/TrasladoDR/@TasaOCuotaDR", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "fractionDigits", n: 6},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_IdDocumento = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/@IdDocumento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 16},
	{tipo: "maxLength", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron022}},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_Serie = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/@Serie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron012}},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_Folio = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/@Folio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron013}},
}}
var tPagos20_Pagos_Pago_DoctoRelacionado_EquivalenciaDR = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/@EquivalenciaDR", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 10},
	{tipo: "minInclusive", valor: "0.0000000001"},
}}
var xsinteger = &tipoXSD{nombre: "xs:integer", prim: "integer", ws: "collapse"}
var tPagos20_Pagos_Pago_DoctoRelacionado_NumParcialidad = &tipoXSD{nombre: "Pagos20 Pagos/Pago/DoctoRelacionado/@NumParcialidad", prim: "integer", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron023}},
}}
var tPagos20_Pagos_Pago_ImpuestosP_TrasladosP_TrasladoP_TasaOCuotaP = &tipoXSD{nombre: "Pagos20 Pagos/Pago/ImpuestosP/TrasladosP/TrasladoP/@TasaOCuotaP", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000000"},
}}
var tPagos20_Pagos_Pago_TipoCambioP = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@TipoCambioP", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tPagos20_Pagos_Pago_NumOperacion = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@NumOperacion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tPagos20_Pagos_Pago_RfcEmisorCtaOrd = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@RfcEmisorCtaOrd", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron024}},
}}
var tPagos20_Pagos_Pago_NomBancoOrdExt = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@NomBancoOrdExt", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tPagos20_Pagos_Pago_CtaOrdenante = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@CtaOrdenante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 10},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron025}},
}}
var tPagos20_Pagos_Pago_CtaBeneficiario = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@CtaBeneficiario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 10},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron025}},
}}
var tCatPagos_cTipoCadenaPago = &tipoXSD{nombre: "c_TipoCadenaPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatPagos_cTipoCadenaPago},
}}
var xsbase64Binary = &tipoXSD{nombre: "xs:base64Binary", prim: "base64Binary", ws: "collapse"}
var tPagos20_Pagos_Pago_CertPago = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@CertPago", prim: "base64Binary", ws: "collapse"}
var tPagos20_Pagos_Pago_CadPago = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@CadPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 8192},
}}
var tPagos20_Pagos_Pago_SelloPago = &tipoXSD{nombre: "Pagos20 Pagos/Pago/@SelloPago", prim: "base64Binary", ws: "collapse"}
var tPagos20_Pagos_Version = &tipoXSD{nombre: "Pagos20 Pagos/@Version", prim: "string", ws: "collapse"}
var tCatNomina_cOrigenRecurso = &tipoXSD{nombre: "c_OrigenRecurso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cOrigenRecurso},
}}
var tTdCFDI_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron026}},
}}
var tNomina12_Nomina_Emisor_RegistroPatronal = &tipoXSD{nombre: "Nomina12 Nomina/Emisor/@RegistroPatronal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tNomina12_Nomina_Receptor_SubContratacion_PorcentajeTiempo = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/SubContratacion/@PorcentajeTiempo", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "maxInclusive", valor: "100.000"},
	{tipo: "pattern", patrones: []patronXSD{patron027}},
}}
var tNomina12_Nomina_Receptor_NumSeguridadSocial = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/@NumSeguridadSocial", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron028}},
}}
var xsdate = &tipoXSD{nombre: "xs:date", prim: "date", ws: "collapse"}
var tTdCFDI_tFecha = &tipoXSD{nombre: "t_Fecha", prim: "date", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron029}},
}}
var tNomina12_Nomina_Receptor_Antiguedad = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/@Antigüedad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron030}},
}}
var tCatNomina_cTipoContrato = &tipoXSD{nombre: "c_TipoContrato", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoContrato},
}}
var tNomina12_Nomina_Receptor_Sindicalizado = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/@Sindicalizado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catNomina12_Nomina_Receptor_Sindicalizado},
}}
var tCatNomina_cTipoJornada = &tipoXSD{nombre: "c_TipoJornada", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoJornada},
}}
var tCatNomina_cTipoRegimen = &tipoXSD{nombre: "c_TipoRegimen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoRegimen},
}}
var tNomina12_Nomina_Receptor_NumEmpleado = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/@NumEmpleado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tNomina12_Nomina_Receptor_Departamento = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/@Departamento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tNomina12_Nomina_Receptor_Puesto = &tipoXSD{nombre: "Nomina12 Nomina/Receptor/@Puesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCatNomina_cRiesgoPuesto = &tipoXSD{nombre: "c_RiesgoPuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cRiesgoPuesto},
}}
var tCatNomina_cPeriodicidadPago = &tipoXSD{nombre: "c_PeriodicidadPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cPeriodicidadPago},
}}
var tCatNomina_cBanco = &tipoXSD{nombre: "c_Banco", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cBanco},
}}
var tTdCFDI_tCuentaBancaria = &tipoXSD{nombre: "t_CuentaBancaria", prim: "integer", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron032}},
}}
var tCatCFDI_cEstado = &tipoXSD{nombre: "c_Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cEstado},
}}
var tNomina12_Nomina_Percepciones_Percepcion_AccionesOTitulos_ValorMercado = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/Percepcion/AccionesOTitulos/@ValorMercado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tNomina12_Nomina_Percepciones_Percepcion_AccionesOTitulos_PrecioAlOtorgarse = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/Percepcion/AccionesOTitulos/@PrecioAlOtorgarse", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var xsint = &tipoXSD{nombre: "xs:int", prim: "int", ws: "collapse"}
var tNomina12_Nomina_Percepciones_Percepcion_HorasExtra_Dias = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/Percepcion/HorasExtra/@Dias", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCatNomina_cTipoHoras = &tipoXSD{nombre: "c_TipoHoras", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoHoras},
}}
var tNomina12_Nomina_Percepciones_Percepcion_HorasExtra_HorasExtra = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/Percepcion/HorasExtra/@HorasExtra", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCatNomina_cTipoPercepcion = &tipoXSD{nombre: "c_TipoPercepcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoPercepcion},
}}
var tNomina12_Nomina_Percepciones_Percepcion_Clave = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/Percepcion/@Clave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron033}},
}}
var tNomina12_Nomina_Percepciones_Percepcion_Concepto = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/Percepcion/@Concepto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tNomina12_Nomina_Percepciones_SeparacionIndemnizacion_NumAniosServicio = &tipoXSD{nombre: "Nomina12 Nomina/Percepciones/SeparacionIndemnizacion/@NumAñosServicio", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0"},
	{tipo: "maxInclusive", valor: "99"},
}}
var tCatNomina_cTipoDeduccion = &tipoXSD{nombre: "c_TipoDeduccion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoDeduccion},
}}
var tNomina12_Nomina_Deducciones_Deduccion_Clave = &tipoXSD{nombre: "Nomina12 Nomina/Deducciones/Deduccion/@Clave", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron033}},
}}
var tNomina12_Nomina_Deducciones_Deduccion_Concepto = &tipoXSD{nombre: "Nomina12 Nomina/Deducciones/Deduccion/@Concepto", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tNomina12_Nomina_OtrosPagos_OtroPago_CompensacionSaldosAFavor_Anio = &tipoXSD{nombre: "Nomina12 Nomina/OtrosPagos/OtroPago/CompensacionSaldosAFavor/@Año", prim: "short", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "2016"},
}}
var tCatNomina_cTipoOtroPago = &tipoXSD{nombre: "c_TipoOtroPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoOtroPago},
}}
var tNomina12_Nomina_OtrosPagos_OtroPago_Clave = &tipoXSD{nombre: "Nomina12 Nomina/OtrosPagos/OtroPago/@Clave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron033}},
}}
var tNomina12_Nomina_OtrosPagos_OtroPago_Concepto = &tipoXSD{nombre: "Nomina12 Nomina/OtrosPagos/OtroPago/@Concepto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tNomina12_Nomina_Incapacidades_Incapacidad_DiasIncapacidad = &tipoXSD{nombre: "Nomina12 Nomina/Incapacidades/Incapacidad/@DiasIncapacidad", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCatNomina_cTipoIncapacidad = &tipoXSD{nombre: "c_TipoIncapacidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoIncapacidad},
}}
var tCatNomina_cTipoNomina = &tipoXSD{nombre: "c_TipoNomina", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatNomina_cTipoNomina},
}}
var tNomina12_Nomina_NumDiasPagados = &tipoXSD{nombre: "Nomina12 Nomina/@NumDiasPagados", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "maxInclusive", valor: "36160.000"},
	{tipo: "fractionDigits", n: 3},
	{tipo: "pattern", patrones: []patronXSD{patron034}},
}}
var tImpLocal_ImpuestosLocales_RetencionesLocales_TasadeRetencion = &tipoXSD{nombre: "ImpLocal ImpuestosLocales/RetencionesLocales/@TasadeRetencion", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tImpLocal_ImpuestosLocales_RetencionesLocales_Importe = &tipoXSD{nombre: "ImpLocal ImpuestosLocales/RetencionesLocales/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tImpLocal_ImpuestosLocales_TrasladosLocales_TasadeTraslado = &tipoXSD{nombre: "ImpLocal ImpuestosLocales/TrasladosLocales/@TasadeTraslado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tImpLocal_ImpuestosLocales_TrasladosLocales_Importe = &tipoXSD{nombre: "ImpLocal ImpuestosLocales/TrasladosLocales/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tImpLocal_ImpuestosLocales_TotaldeRetenciones = &tipoXSD{nombre: "ImpLocal ImpuestosLocales/@TotaldeRetenciones", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tImpLocal_ImpuestosLocales_TotaldeTraslados = &tipoXSD{nombre: "ImpLocal ImpuestosLocales/@TotaldeTraslados", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tCatCartaPorte_cRegimenAduanero = &tipoXSD{nombre: "c_RegimenAduanero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cRegimenAduanero},
}}
var tTdCFDI_tDescrip100 = &tipoXSD{nombre: "t_Descrip100", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tTdCFDI_tNumeroDomicilio = &tipoXSD{nombre: "t_NumeroDomicilio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 55},
	{tipo: "pattern", patrones: []patronXSD{patron035}},
}}
var tTdCFDI_tDescrip120 = &tipoXSD{nombre: "t_Descrip120", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 120},
	{tipo: "pattern", patrones: []patronXSD{patron036}},
}}
var tTdCFDI_tReferencia = &tipoXSD{nombre: "t_Referencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 250},
	{tipo: "pattern", patrones: []patronXSD{patron037}},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@TipoUbicacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_IDUbicacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@IDUbicacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron039}},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NombreRemitenteDestinatario = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@NombreRemitenteDestinatario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NumRegIdTrib = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCatCartaPorte_cEstaciones = &tipoXSD{nombre: "c_Estaciones", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cEstaciones},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NombreEstacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@NombreEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@NavegacionTrafico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico},
}}
var tCatCartaPorte_cTipoEstacion = &tipoXSD{nombre: "c_TipoEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoEstacion},
}}
var tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Ubicaciones/Ubicacion/@DistanciaRecorrida", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCatCartaPorte_cDocumentoAduanero = &tipoXSD{nombre: "c_DocumentoAduanero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cDocumentoAduanero},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_NumPedimento = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/DocumentacionAduanera/@NumPedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "maxLength", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_IdentDocAduanero = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/DocumentacionAduanera/@IdentDocAduanero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron042}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_NumeroGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@NumeroGuiaIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron043}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_DescripGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@DescripGuiaIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_PesoGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@PesoGuiaIdentificacion", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron044}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDDestino", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron045}},
}}
var tCatCartaPorte_cCveTransporte = &tipoXSD{nombre: "c_CveTransporte", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cCveTransporte},
}}
var tCatCartaPorte_cClaveUnidadPeso = &tipoXSD{nombre: "c_ClaveUnidadPeso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cClaveUnidadPeso},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoNeto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoTara", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@NumPiezas", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCatCartaPorte_cClaveProdServCP = &tipoXSD{nombre: "c_ClaveProdServCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cClaveProdServCP},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_ClaveSTCC = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@ClaveSTCC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 7},
	{tipo: "pattern", patrones: []patronXSD{patron046}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_Descripcion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_Cantidad = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_Unidad = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_Dimensiones = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@Dimensiones", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron047}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_MaterialPeligroso = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@MaterialPeligroso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
}}
var tCatCartaPorte_cMaterialPeligroso = &tipoXSD{nombre: "c_MaterialPeligroso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cMaterialPeligroso},
}}
var tCatCartaPorte_cTipoEmbalaje = &tipoXSD{nombre: "c_TipoEmbalaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoEmbalaje},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DescripEmbalaje = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DescripEmbalaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCatCartaPorte_cSectorCOFEPRIS = &tipoXSD{nombre: "c_SectorCOFEPRIS", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cSectorCOFEPRIS},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_NombreIngredienteActivo = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@NombreIngredienteActivo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_NomQuimico = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@NomQuimico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron042}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DenominacionGenericaProd = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DenominacionGenericaProd", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DenominacionDistintivaProd = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DenominacionDistintivaProd", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_Fabricante = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@Fabricante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 240},
	{tipo: "pattern", patrones: []patronXSD{patron048}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_LoteMedicamento = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@LoteMedicamento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
	{tipo: "pattern", patrones: []patronXSD{patron049}},
}}
var tCatCartaPorte_cFormaFarmaceutica = &tipoXSD{nombre: "c_FormaFarmaceutica", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cFormaFarmaceutica},
}}
var tCatCartaPorte_cCondicionesEspeciales = &tipoXSD{nombre: "c_CondicionesEspeciales", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cCondicionesEspeciales},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_RegistroSanitarioFolioAutorizacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@RegistroSanitarioFolioAutorizacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_PermisoImportacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@PermisoImportacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 6},
	{tipo: "pattern", patrones: []patronXSD{patron050}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_FolioImpoVUCEM = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@FolioImpoVUCEM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron012}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_NumCAS = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@NumCAS", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_RazonSocialEmpImp = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@RazonSocialEmpImp", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 80},
	{tipo: "pattern", patrones: []patronXSD{patron051}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_NumRegSanPlagCOFEPRIS = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@NumRegSanPlagCOFEPRIS", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 60},
	{tipo: "pattern", patrones: []patronXSD{patron052}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DatosFabricante = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DatosFabricante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 600},
	{tipo: "pattern", patrones: []patronXSD{patron053}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DatosFormulador = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DatosFormulador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 600},
	{tipo: "pattern", patrones: []patronXSD{patron053}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DatosMaquilador = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DatosMaquilador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 600},
	{tipo: "pattern", patrones: []patronXSD{patron053}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_UsoAutorizado = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@UsoAutorizado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_PesoEnKg = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@PesoEnKg", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
}}
var tCatComExt_cFraccionArancelaria = &tipoXSD{nombre: "c_FraccionArancelaria", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatComExt_cFraccionArancelaria},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_UUIDComercioExt = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@UUIDComercioExt", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tCatCartaPorte_cTipoMateria = &tipoXSD{nombre: "c_TipoMateria", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoMateria},
}}
var tCartaPorte31_CartaPorte_Mercancias_Mercancia_DescripcionMateria = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Mercancia/@DescripcionMateria", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCatCartaPorte_cConfigAutotransporte = &tipoXSD{nombre: "c_ConfigAutotransporte", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cConfigAutotransporte},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PesoBrutoVehicular = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@PesoBrutoVehicular", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "fractionDigits", n: 2},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PlacaVM = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@PlacaVM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_AnioModeloVM = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@AnioModeloVM", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraRespCivil = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraRespCivil", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaRespCivil = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaRespCivil", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraMedAmbiente = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraMedAmbiente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaMedAmbiente = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaMedAmbiente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraCarga = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaCarga = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCatCartaPorte_cSubTipoRem = &tipoXSD{nombre: "c_SubTipoRem", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cSubTipoRem},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque_Placa = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/Remolques/Remolque/@Placa", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCatCartaPorte_cTipoPermiso = &tipoXSD{nombre: "c_TipoPermiso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoPermiso},
}}
var tCartaPorte31_CartaPorte_Mercancias_Autotransporte_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/Autotransporte/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_RemolquesCCP_RemolqueCCP_PlacaCCP = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/RemolquesCCP/RemolqueCCP/@PlacaCCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCatCartaPorte_cContenedorMaritimo = &tipoXSD{nombre: "c_ContenedorMaritimo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cContenedorMaritimo},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@MatriculaContenedor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 11},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron058}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@NumPrecinto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron059}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_IdCCPRelacionado = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@IdCCPRelacionado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron060}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_PlacaVMCCP = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@PlacaVMCCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCatCartaPorte_cConfigMaritima = &tipoXSD{nombre: "c_ConfigMaritima", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cConfigMaritima},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Matricula = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@Matricula", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 7},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron061}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NumeroOMI", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron062}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@AnioEmbarcacion", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NombreEmbarc", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@UnidadesDeArqBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCatCartaPorte_cClaveTipoCarga = &tipoXSD{nombre: "c_ClaveTipoCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cClaveTipoCarga},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Eslora = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@Eslora", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Manga = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@Manga", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Calado = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@Calado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Puntal = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@Puntal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "fractionDigits", n: 2},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@LineaNaviera", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NombreAgenteNaviero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCatCartaPorte_cNumAutorizacionNaviero = &tipoXSD{nombre: "c_NumAutorizacionNaviero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cNumAutorizacionNaviero},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumViaje = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NumViaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@NumConocEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_PermisoTempNavegacion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteMaritimo/@PermisoTempNavegacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 19},
	{tipo: "pattern", patrones: []patronXSD{patron063}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@MatriculaAeronave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 11},
	{tipo: "pattern", patrones: []patronXSD{patron064}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NombreAseg = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumeroGuia = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@NumeroGuia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron065}},
}}
var tCatCartaPorte_cCodigoTransporteAereo = &tipoXSD{nombre: "c_CodigoTransporteAereo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cCodigoTransporteAereo},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@NumRegIdTribEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteAereo/@NombreEmbarcador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCatCartaPorte_cDerechosDePaso = &tipoXSD{nombre: "c_DerechosDePaso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cDerechosDePaso},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/DerechosDePaso/@KilometrajePagado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCatCartaPorte_cContenedor = &tipoXSD{nombre: "c_Contenedor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cContenedor},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoContenedorVacio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoNetoMercancia", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCatCartaPorte_cTipoCarro = &tipoXSD{nombre: "c_TipoCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoCarro},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/Carro/@MatriculaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron066}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/Carro/@GuiaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/Carro/@ToneladasNetasCarro", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCatCartaPorte_cTipoDeServicio = &tipoXSD{nombre: "c_TipoDeServicio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoDeServicio},
}}
var tCatCartaPorte_cTipoDeTrafico = &tipoXSD{nombre: "c_TipoDeTrafico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cTipoDeTrafico},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/TransporteFerroviario/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte31_CartaPorte_Mercancias_PesoBrutoTotal = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/@PesoBrutoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_PesoNetoTotal = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/@PesoNetoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte31_CartaPorte_Mercancias_NumTotalMercancias = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/@NumTotalMercancias", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCartaPorte31_CartaPorte_Mercancias_LogisticaInversaRecoleccionDevolucion = &tipoXSD{nombre: "CartaPorte31 CartaPorte/Mercancias/@LogisticaInversaRecoleccionDevolucion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_Mercancias_LogisticaInversaRecoleccionDevolucion},
}}
var tCatCartaPorte_cParteTransporte = &tipoXSD{nombre: "c_ParteTransporte", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cParteTransporte},
}}
var tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte31 CartaPorte/FiguraTransporte/TiposFigura/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte31 CartaPorte/FiguraTransporte/TiposFigura/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCatCartaPorte_cFiguraTransporte = &tipoXSD{nombre: "c_FiguraTransporte", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cFiguraTransporte},
}}
var tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_RFCFigura = &tipoXSD{nombre: "CartaPorte31 CartaPorte/FiguraTransporte/TiposFigura/@RFCFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_NumLicencia = &tipoXSD{nombre: "CartaPorte31 CartaPorte/FiguraTransporte/TiposFigura/@NumLicencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 16},
	{tipo: "pattern", patrones: []patronXSD{patron067}},
}}
var tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_NombreFigura = &tipoXSD{nombre: "CartaPorte31 CartaPorte/FiguraTransporte/TiposFigura/@NombreFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_NumRegIdTribFigura = &tipoXSD{nombre: "CartaPorte31 CartaPorte/FiguraTransporte/TiposFigura/@NumRegIdTribFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte31_CartaPorte_IdCCP = &tipoXSD{nombre: "CartaPorte31 CartaPorte/@IdCCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron060}},
}}
var tCartaPorte31_CartaPorte_TranspInternac = &tipoXSD{nombre: "CartaPorte31 CartaPorte/@TranspInternac", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_TranspInternac},
}}
var tCartaPorte31_CartaPorte_EntradaSalidaMerc = &tipoXSD{nombre: "CartaPorte31 CartaPorte/@EntradaSalidaMerc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_EntradaSalidaMerc},
}}
var tCartaPorte31_CartaPorte_TotalDistRec = &tipoXSD{nombre: "CartaPorte31 CartaPorte/@TotalDistRec", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCartaPorte31_CartaPorte_RegistroISTMO = &tipoXSD{nombre: "CartaPorte31 CartaPorte/@RegistroISTMO", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte31_CartaPorte_RegistroISTMO},
}}
var tCatCartaPorte_cRegistroISTMO = &tipoXSD{nombre: "c_RegistroISTMO", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cRegistroISTMO},
}}
var tPagos10_Pagos_Pago_DoctoRelacionado_IdDocumento = &tipoXSD{nombre: "Pagos10 Pagos/Pago/DoctoRelacionado/@IdDocumento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 16},
	{tipo: "maxLength", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron022}},
}}
var tPagos10_Pagos_Pago_DoctoRelacionado_Serie = &tipoXSD{nombre: "Pagos10 Pagos/Pago/DoctoRelacionado/@Serie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron012}},
}}
var tPagos10_Pagos_Pago_DoctoRelacionado_Folio = &tipoXSD{nombre: "Pagos10 Pagos/Pago/DoctoRelacionado/@Folio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron013}},
}}
var tPagos10_Pagos_Pago_DoctoRelacionado_TipoCambioDR = &tipoXSD{nombre: "Pagos10 Pagos/Pago/DoctoRelacionado/@TipoCambioDR", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tPagos10_Pagos_Pago_DoctoRelacionado_NumParcialidad = &tipoXSD{nombre: "Pagos10 Pagos/Pago/DoctoRelacionado/@NumParcialidad", prim: "integer", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron023}},
}}
var tPagos10_Pagos_Pago_Impuestos_Traslados_Traslado_TasaOCuota = &tipoXSD{nombre: "Pagos10 Pagos/Pago/Impuestos/Traslados/Traslado/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000000"},
}}
var tPagos10_Pagos_Pago_TipoCambioP = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@TipoCambioP", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tPagos10_Pagos_Pago_NumOperacion = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@NumOperacion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tPagos10_Pagos_Pago_RfcEmisorCtaOrd = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@RfcEmisorCtaOrd", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron024}},
}}
var tPagos10_Pagos_Pago_NomBancoOrdExt = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@NomBancoOrdExt", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tPagos10_Pagos_Pago_CtaOrdenante = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@CtaOrdenante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 10},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron025}},
}}
var tPagos10_Pagos_Pago_CtaBeneficiario = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@CtaBeneficiario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 10},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron025}},
}}
var tPagos10_Pagos_Pago_CertPago = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@CertPago", prim: "base64Binary", ws: "collapse"}
var tPagos10_Pagos_Pago_CadPago = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@CadPago", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 8192},
}}
var tPagos10_Pagos_Pago_SelloPago = &tipoXSD{nombre: "Pagos10 Pagos/Pago/@SelloPago", prim: "base64Binary", ws: "collapse"}
var tPagos10_Pagos_Version = &tipoXSD{nombre: "Pagos10 Pagos/@Version", prim: "string", ws: "collapse"}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@TipoUbicacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_IDUbicacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@IDUbicacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron039}},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NombreRemitenteDestinatario = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@NombreRemitenteDestinatario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NumRegIdTrib = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NombreEstacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@NombreEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@NavegacionTrafico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico},
}}
var tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Ubicaciones/Ubicacion/@DistanciaRecorrida", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_NumPedimento = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/DocumentacionAduanera/@NumPedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "maxLength", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_IdentDocAduanero = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/DocumentacionAduanera/@IdentDocAduanero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron042}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_NumeroGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@NumeroGuiaIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron043}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_DescripGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@DescripGuiaIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_PesoGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@PesoGuiaIdentificacion", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron044}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDDestino", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron045}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoNeto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoTara", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@NumPiezas", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_ClaveSTCC = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@ClaveSTCC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 7},
	{tipo: "pattern", patrones: []patronXSD{patron046}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_Descripcion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_Cantidad = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_Unidad = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_Dimensiones = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@Dimensiones", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron047}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_MaterialPeligroso = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@MaterialPeligroso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DescripEmbalaje = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DescripEmbalaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_NombreIngredienteActivo = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@NombreIngredienteActivo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_NomQuimico = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@NomQuimico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
	{tipo: "pattern", patrones: []patronXSD{patron042}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DenominacionGenericaProd = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DenominacionGenericaProd", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DenominacionDistintivaProd = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DenominacionDistintivaProd", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_Fabricante = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@Fabricante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 240},
	{tipo: "pattern", patrones: []patronXSD{patron048}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_LoteMedicamento = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@LoteMedicamento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
	{tipo: "pattern", patrones: []patronXSD{patron049}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_RegistroSanitarioFolioAutorizacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@RegistroSanitarioFolioAutorizacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_PermisoImportacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@PermisoImportacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 6},
	{tipo: "pattern", patrones: []patronXSD{patron050}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_FolioImpoVUCEM = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@FolioImpoVUCEM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron012}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_NumCAS = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@NumCAS", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_RazonSocialEmpImp = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@RazonSocialEmpImp", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 80},
	{tipo: "pattern", patrones: []patronXSD{patron051}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_NumRegSanPlagCOFEPRIS = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@NumRegSanPlagCOFEPRIS", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 60},
	{tipo: "pattern", patrones: []patronXSD{patron052}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DatosFabricante = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DatosFabricante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 600},
	{tipo: "pattern", patrones: []patronXSD{patron053}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DatosFormulador = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DatosFormulador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 600},
	{tipo: "pattern", patrones: []patronXSD{patron053}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DatosMaquilador = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DatosMaquilador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 600},
	{tipo: "pattern", patrones: []patronXSD{patron053}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_UsoAutorizado = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@UsoAutorizado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_PesoEnKg = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@PesoEnKg", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_UUIDComercioExt = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@UUIDComercioExt", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Mercancia_DescripcionMateria = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Mercancia/@DescripcionMateria", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PesoBrutoVehicular = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@PesoBrutoVehicular", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "fractionDigits", n: 2},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PlacaVM = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@PlacaVM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_AnioModeloVM = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@AnioModeloVM", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraRespCivil = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraRespCivil", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaRespCivil = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaRespCivil", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraMedAmbiente = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraMedAmbiente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaMedAmbiente = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaMedAmbiente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraCarga = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaCarga = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque_Placa = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/Remolques/Remolque/@Placa", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte30_CartaPorte_Mercancias_Autotransporte_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/Autotransporte/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@MatriculaContenedor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 11},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron058}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@NumPrecinto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron059}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_IdCCPRelacionado = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@IdCCPRelacionado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron060}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_PlacaVMCCP = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@PlacaVMCCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_RemolquesCCP_RemolqueCCP_PlacaCCP = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/RemolquesCCP/RemolqueCCP/@PlacaCCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Matricula = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@Matricula", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 7},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron061}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NumeroOMI", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron062}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@AnioEmbarcacion", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NombreEmbarc", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@UnidadesDeArqBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Eslora = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@Eslora", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Manga = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@Manga", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Calado = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@Calado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Puntal = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@Puntal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "fractionDigits", n: 2},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@LineaNaviera", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NombreAgenteNaviero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumViaje = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NumViaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@NumConocEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_PermisoTempNavegacion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteMaritimo/@PermisoTempNavegacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 19},
	{tipo: "pattern", patrones: []patronXSD{patron063}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@MatriculaAeronave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 11},
	{tipo: "pattern", patrones: []patronXSD{patron064}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NombreAseg = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumeroGuia = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@NumeroGuia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron065}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@NumRegIdTribEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteAereo/@NombreEmbarcador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/DerechosDePaso/@KilometrajePagado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoContenedorVacio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoNetoMercancia", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/Carro/@MatriculaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron066}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/Carro/@GuiaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/Carro/@ToneladasNetasCarro", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/TransporteFerroviario/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte30_CartaPorte_Mercancias_PesoBrutoTotal = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/@PesoBrutoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_PesoNetoTotal = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/@PesoNetoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte30_CartaPorte_Mercancias_NumTotalMercancias = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/@NumTotalMercancias", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCartaPorte30_CartaPorte_Mercancias_LogisticaInversaRecoleccionDevolucion = &tipoXSD{nombre: "CartaPorte30 CartaPorte/Mercancias/@LogisticaInversaRecoleccionDevolucion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_Mercancias_LogisticaInversaRecoleccionDevolucion},
}}
var tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte30 CartaPorte/FiguraTransporte/TiposFigura/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte30 CartaPorte/FiguraTransporte/TiposFigura/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_RFCFigura = &tipoXSD{nombre: "CartaPorte30 CartaPorte/FiguraTransporte/TiposFigura/@RFCFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_NumLicencia = &tipoXSD{nombre: "CartaPorte30 CartaPorte/FiguraTransporte/TiposFigura/@NumLicencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 16},
	{tipo: "pattern", patrones: []patronXSD{patron067}},
}}
var tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_NombreFigura = &tipoXSD{nombre: "CartaPorte30 CartaPorte/FiguraTransporte/TiposFigura/@NombreFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_NumRegIdTribFigura = &tipoXSD{nombre: "CartaPorte30 CartaPorte/FiguraTransporte/TiposFigura/@NumRegIdTribFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte30_CartaPorte_IdCCP = &tipoXSD{nombre: "CartaPorte30 CartaPorte/@IdCCP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron060}},
}}
var tCartaPorte30_CartaPorte_TranspInternac = &tipoXSD{nombre: "CartaPorte30 CartaPorte/@TranspInternac", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_TranspInternac},
}}
var tCartaPorte30_CartaPorte_EntradaSalidaMerc = &tipoXSD{nombre: "CartaPorte30 CartaPorte/@EntradaSalidaMerc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_EntradaSalidaMerc},
}}
var tCartaPorte30_CartaPorte_TotalDistRec = &tipoXSD{nombre: "CartaPorte30 CartaPorte/@TotalDistRec", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCartaPorte30_CartaPorte_RegistroISTMO = &tipoXSD{nombre: "CartaPorte30 CartaPorte/@RegistroISTMO", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte30_CartaPorte_RegistroISTMO},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@TipoUbicacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_IDUbicacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@IDUbicacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron039}},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NombreRemitenteDestinatario = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@NombreRemitenteDestinatario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NumRegIdTrib = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NombreEstacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@NombreEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@NavegacionTrafico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico},
}}
var tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Ubicaciones/Ubicacion/@DistanciaRecorrida", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_Pedimentos_Pedimento = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/Pedimentos/@Pedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_NumeroGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@NumeroGuiaIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron043}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_DescripGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@DescripGuiaIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_PesoGuiaIdentificacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/GuiasIdentificacion/@PesoGuiaIdentificacion", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron044}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDDestino", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron045}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoNeto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoTara", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@NumPiezas", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_ClaveSTCC = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@ClaveSTCC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 7},
	{tipo: "pattern", patrones: []patronXSD{patron046}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_Descripcion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_Cantidad = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_Unidad = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_Dimensiones = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@Dimensiones", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron047}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_MaterialPeligroso = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@MaterialPeligroso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte20_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_DescripEmbalaje = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@DescripEmbalaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_PesoEnKg = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@PesoEnKg", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
}}
var tCartaPorte20_CartaPorte_Mercancias_Mercancia_UUIDComercioExt = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Mercancia/@UUIDComercioExt", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PlacaVM = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@PlacaVM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_AnioModeloVM = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/IdentificacionVehicular/@AnioModeloVM", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraRespCivil = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraRespCivil", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaRespCivil = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaRespCivil", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraMedAmbiente = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraMedAmbiente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaMedAmbiente = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaMedAmbiente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraCarga = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Seguros/@AseguraCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaCarga = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Seguros/@PolizaCarga", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque_Placa = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/Remolques/Remolque/@Placa", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron054}},
}}
var tCartaPorte20_CartaPorte_Mercancias_Autotransporte_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/Autotransporte/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@MatriculaContenedor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 11},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron058}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@NumPrecinto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron059}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Matricula = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@Matricula", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 7},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron061}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NumeroOMI", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron062}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@AnioEmbarcacion", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NombreEmbarc", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@UnidadesDeArqBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumCertITC = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NumCertITC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron068}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Eslora = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@Eslora", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Manga = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@Manga", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Calado = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@Calado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@LineaNaviera", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NombreAgenteNaviero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumViaje = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NumViaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteMaritimo/@NumConocEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@MatriculaAeronave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 10},
	{tipo: "pattern", patrones: []patronXSD{patron069}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NombreAseg = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumeroGuia = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@NumeroGuia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron065}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@NumRegIdTribEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteAereo/@NombreEmbarcador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/DerechosDePaso/@KilometrajePagado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoContenedorVacio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoNetoMercancia", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/Carro/@MatriculaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron066}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/Carro/@GuiaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/Carro/@ToneladasNetasCarro", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/TransporteFerroviario/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte20_CartaPorte_Mercancias_PesoBrutoTotal = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/@PesoBrutoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_PesoNetoTotal = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/@PesoNetoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte20_CartaPorte_Mercancias_NumTotalMercancias = &tipoXSD{nombre: "CartaPorte20 CartaPorte/Mercancias/@NumTotalMercancias", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte20 CartaPorte/FiguraTransporte/TiposFigura/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte20 CartaPorte/FiguraTransporte/TiposFigura/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_RFCFigura = &tipoXSD{nombre: "CartaPorte20 CartaPorte/FiguraTransporte/TiposFigura/@RFCFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_NumLicencia = &tipoXSD{nombre: "CartaPorte20 CartaPorte/FiguraTransporte/TiposFigura/@NumLicencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 16},
	{tipo: "pattern", patrones: []patronXSD{patron067}},
}}
var tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_NombreFigura = &tipoXSD{nombre: "CartaPorte20 CartaPorte/FiguraTransporte/TiposFigura/@NombreFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_NumRegIdTribFigura = &tipoXSD{nombre: "CartaPorte20 CartaPorte/FiguraTransporte/TiposFigura/@NumRegIdTribFigura", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte20_CartaPorte_TranspInternac = &tipoXSD{nombre: "CartaPorte20 CartaPorte/@TranspInternac", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte20_CartaPorte_TranspInternac},
}}
var tCartaPorte20_CartaPorte_EntradaSalidaMerc = &tipoXSD{nombre: "CartaPorte20 CartaPorte/@EntradaSalidaMerc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte20_CartaPorte_EntradaSalidaMerc},
}}
var tCartaPorte20_CartaPorte_TotalDistRec = &tipoXSD{nombre: "CartaPorte20 CartaPorte/@TotalDistRec", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCatCFDI_cColonia = &tipoXSD{nombre: "c_Colonia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cColonia},
}}
var tCatCFDI_cLocalidad = &tipoXSD{nombre: "c_Localidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cLocalidad},
}}
var tCatCFDI_cMunicipio = &tipoXSD{nombre: "c_Municipio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCFDI_cMunicipio},
}}
var tComExt20_ComercioExterior_Propietario_NumRegIdTrib = &tipoXSD{nombre: "ComExt20 ComercioExterior/Propietario/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt20_ComercioExterior_Receptor_Domicilio_Estado = &tipoXSD{nombre: "ComExt20 ComercioExterior/Receptor/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tComExt20_ComercioExterior_Receptor_Domicilio_CodigoPostal = &tipoXSD{nombre: "ComExt20 ComercioExterior/Receptor/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tComExt20_ComercioExterior_Receptor_NumRegIdTrib = &tipoXSD{nombre: "ComExt20 ComercioExterior/Receptor/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt20_ComercioExterior_Destinatario_Domicilio_Estado = &tipoXSD{nombre: "ComExt20 ComercioExterior/Destinatario/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tComExt20_ComercioExterior_Destinatario_Domicilio_CodigoPostal = &tipoXSD{nombre: "ComExt20 ComercioExterior/Destinatario/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tComExt20_ComercioExterior_Destinatario_NumRegIdTrib = &tipoXSD{nombre: "ComExt20 ComercioExterior/Destinatario/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt20_ComercioExterior_Destinatario_Nombre = &tipoXSD{nombre: "ComExt20 ComercioExterior/Destinatario/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Marca = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@Marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 35},
	{tipo: "pattern", patrones: []patronXSD{patron070}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Modelo = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@Modelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 80},
	{tipo: "pattern", patrones: []patronXSD{patron051}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_SubModelo = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@SubModelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_NumeroSerie = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@NumeroSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron013}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_NoIdentificacion = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_CantidadAduana = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/@CantidadAduana", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "pattern", patrones: []patronXSD{patron071}},
}}
var tCatComExt_cUnidadAduana = &tipoXSD{nombre: "c_UnidadAduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatComExt_cUnidadAduana},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_ValorUnitarioAduana = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/@ValorUnitarioAduana", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.000000"},
	{tipo: "pattern", patrones: []patronXSD{patron072}},
}}
var tComExt20_ComercioExterior_Mercancias_Mercancia_ValorDolares = &tipoXSD{nombre: "ComExt20 ComercioExterior/Mercancias/Mercancia/@ValorDolares", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.0000"},
	{tipo: "pattern", patrones: []patronXSD{patron073}},
}}
var tComExt20_ComercioExterior_Version = &tipoXSD{nombre: "ComExt20 ComercioExterior/@Version", prim: "string", ws: "collapse"}
var tCatComExt_cMotivoTraslado = &tipoXSD{nombre: "c_MotivoTraslado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatComExt_cMotivoTraslado},
}}
var tCatComExt_cClavePedimento = &tipoXSD{nombre: "c_ClavePedimento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatComExt_cClavePedimento},
}}
var tComExt20_ComercioExterior_CertificadoOrigen = &tipoXSD{nombre: "ComExt20 ComercioExterior/@CertificadoOrigen", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt20_ComercioExterior_CertificadoOrigen},
}}
var tComExt20_ComercioExterior_NumCertificadoOrigen = &tipoXSD{nombre: "ComExt20 ComercioExterior/@NumCertificadoOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tComExt20_ComercioExterior_NumeroExportadorConfiable = &tipoXSD{nombre: "ComExt20 ComercioExterior/@NumeroExportadorConfiable", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCatComExt_cINCOTERM = &tipoXSD{nombre: "c_INCOTERM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatComExt_cINCOTERM},
}}
var tComExt20_ComercioExterior_Observaciones = &tipoXSD{nombre: "ComExt20 ComercioExterior/@Observaciones", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tTdCFDI_tTipoCambio = &tipoXSD{nombre: "t_TipoCambio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "fractionDigits", n: 6},
	{tipo: "pattern", patrones: []patronXSD{patron006}},
}}
var tComExt11_ComercioExterior_Propietario_NumRegIdTrib = &tipoXSD{nombre: "ComExt11 ComercioExterior/Propietario/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt11_ComercioExterior_Receptor_Domicilio_Estado = &tipoXSD{nombre: "ComExt11 ComercioExterior/Receptor/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tComExt11_ComercioExterior_Receptor_Domicilio_CodigoPostal = &tipoXSD{nombre: "ComExt11 ComercioExterior/Receptor/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tComExt11_ComercioExterior_Receptor_NumRegIdTrib = &tipoXSD{nombre: "ComExt11 ComercioExterior/Receptor/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt11_ComercioExterior_Destinatario_Domicilio_Estado = &tipoXSD{nombre: "ComExt11 ComercioExterior/Destinatario/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tComExt11_ComercioExterior_Destinatario_Domicilio_CodigoPostal = &tipoXSD{nombre: "ComExt11 ComercioExterior/Destinatario/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tComExt11_ComercioExterior_Destinatario_NumRegIdTrib = &tipoXSD{nombre: "ComExt11 ComercioExterior/Destinatario/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt11_ComercioExterior_Destinatario_Nombre = &tipoXSD{nombre: "ComExt11 ComercioExterior/Destinatario/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Marca = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@Marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 35},
	{tipo: "pattern", patrones: []patronXSD{patron070}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Modelo = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@Modelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 80},
	{tipo: "pattern", patrones: []patronXSD{patron051}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_SubModelo = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@SubModelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_NumeroSerie = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@NumeroSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron013}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_NoIdentificacion = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_CantidadAduana = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/@CantidadAduana", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "pattern", patrones: []patronXSD{patron071}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_ValorUnitarioAduana = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/@ValorUnitarioAduana", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "pattern", patrones: []patronXSD{patron074}},
}}
var tComExt11_ComercioExterior_Mercancias_Mercancia_ValorDolares = &tipoXSD{nombre: "ComExt11 ComercioExterior/Mercancias/Mercancia/@ValorDolares", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "pattern", patrones: []patronXSD{patron074}},
}}
var tComExt11_ComercioExterior_Version = &tipoXSD{nombre: "ComExt11 ComercioExterior/@Version", prim: "string", ws: "collapse"}
var tCatComExt_cTipoOperacion = &tipoXSD{nombre: "c_TipoOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatComExt_cTipoOperacion},
}}
var tComExt11_ComercioExterior_CertificadoOrigen = &tipoXSD{nombre: "ComExt11 ComercioExterior/@CertificadoOrigen", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt11_ComercioExterior_CertificadoOrigen},
}}
var tComExt11_ComercioExterior_NumCertificadoOrigen = &tipoXSD{nombre: "ComExt11 ComercioExterior/@NumCertificadoOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tComExt11_ComercioExterior_NumeroExportadorConfiable = &tipoXSD{nombre: "ComExt11 ComercioExterior/@NumeroExportadorConfiable", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tComExt11_ComercioExterior_Subdivision = &tipoXSD{nombre: "ComExt11 ComercioExterior/@Subdivision", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt11_ComercioExterior_Subdivision},
}}
var tComExt11_ComercioExterior_Observaciones = &tipoXSD{nombre: "ComExt11 ComercioExterior/@Observaciones", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
	{tipo: "pattern", patrones: []patronXSD{patron003}},
}}
var tCatHidroYPetro_cTipoPermiso = &tipoXSD{nombre: "c_TipoPermiso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidroYPetro_cTipoPermiso},
}}
var tHidroYPetro_HidroYPetro_NumeroPermiso = &tipoXSD{nombre: "HidroYPetro HidroYPetro/@NumeroPermiso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 15},
	{tipo: "maxLength", n: 35},
}}
var tCatHidroYPetro_cClaveHYP = &tipoXSD{nombre: "c_ClaveHYP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidroYPetro_cClaveHYP},
}}
var tCatHidroYPetro_cSubProductoHYP = &tipoXSD{nombre: "c_SubProductoHYP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidroYPetro_cSubProductoHYP},
}}
var tRegistroFiscal_CFDIRegistroFiscal_Folio = &tipoXSD{nombre: "RegistroFiscal CFDIRegistroFiscal/@Folio", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron075}},
}}
var tDivisas_Divisas_tipoOperacion = &tipoXSD{nombre: "Divisas Divisas/@tipoOperacion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catDivisas_Divisas_tipoOperacion},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Impuesto = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@Impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Impuesto},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_TasaOCuota = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@TasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Importe = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Identificador = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Identificador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ClaveEstacion = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@ClaveEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Cantidad = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
	{tipo: "pattern", patrones: []patronXSD{patron077}},
}}
var tCatCombustible_cClaveTipoCombustible = &tipoXSD{nombre: "c_ClaveTipoCombustible", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCombustible_cClaveTipoCombustible},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Unidad = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_NombreCombustible = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@NombreCombustible", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_FolioOperacion = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@FolioOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ValorUnitario = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@ValorUnitario", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
	{tipo: "pattern", patrones: []patronXSD{patron077}},
}}
var tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Importe = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tECC12_EstadoDeCuentaCombustible_Version = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/@Version", prim: "string", ws: "collapse"}
var tECC12_EstadoDeCuentaCombustible_TipoOperacion = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/@TipoOperacion", prim: "string", ws: "collapse"}
var tECC12_EstadoDeCuentaCombustible_NumeroDeCuenta = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/@NumeroDeCuenta", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tECC12_EstadoDeCuentaCombustible_SubTotal = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/@SubTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tECC12_EstadoDeCuentaCombustible_Total = &tipoXSD{nombre: "ECC12 EstadoDeCuentaCombustible/@Total", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_calle = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@calle", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_noExterior = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@noExterior", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_noInterior = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@noInterior", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_colonia = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@colonia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_localidad = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@localidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_referencia = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@referencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_municipio = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@municipio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_estado = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_pais = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@pais", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_codigoPostal = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionFiscalTercero/@codigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 5},
}}
var tTerceros11_PorCuentadeTerceros_InformacionAduanera_numero = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_InformacionAduanera_fecha = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tTerceros11_PorCuentadeTerceros_InformacionAduanera_aduana = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera_numero = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera_fecha = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera_aduana = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_Parte_cantidad = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/@cantidad", prim: "decimal", ws: "collapse"}
var tTerceros11_PorCuentadeTerceros_Parte_unidad = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/@unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_Parte_noIdentificacion = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/@noIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_Parte_descripcion = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Parte/@descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tTerceros11_PorCuentadeTerceros_CuentaPredial_numero = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/CuentaPredial/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTerceros11_PorCuentadeTerceros_Impuestos_Retenciones_Retencion_impuesto = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Impuestos/Retenciones/Retencion/@impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catTerceros11_PorCuentadeTerceros_Impuestos_Retenciones_Retencion_impuesto},
}}
var tTerceros11_PorCuentadeTerceros_Impuestos_Traslados_Traslado_impuesto = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/Impuestos/Traslados/Traslado/@impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catTerceros11_PorCuentadeTerceros_Impuestos_Traslados_Traslado_impuesto},
}}
var tTerceros11_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
}}
var tTerceros11_PorCuentadeTerceros_nombre = &tipoXSD{nombre: "Terceros11 PorCuentadeTerceros/@nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_InformacionAduanera_numero = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_InformacionAduanera_fecha = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tVentaVehiculos11_VentaVehiculos_InformacionAduanera_aduana = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera_numero = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera_fecha = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera_aduana = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_Parte_cantidad = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/@cantidad", prim: "decimal", ws: "collapse"}
var tVentaVehiculos11_VentaVehiculos_Parte_unidad = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/@unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_Parte_noIdentificacion = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/@noIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_Parte_descripcion = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/Parte/@descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tVentaVehiculos11_VentaVehiculos_ClaveVehicular = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/@ClaveVehicular", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos11_VentaVehiculos_Niv = &tipoXSD{nombre: "VentaVehiculos11 VentaVehiculos/@Niv", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tAerolineas_Aerolineas_OtrosCargos_Cargo_CodigoCargo = &tipoXSD{nombre: "Aerolineas Aerolineas/OtrosCargos/Cargo/@CodigoCargo", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 8},
}}
var tAerolineas_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tLeyendasFisc_LeyendasFiscales_Leyenda_disposicionFiscal = &tipoXSD{nombre: "LeyendasFisc LeyendasFiscales/Leyenda/@disposicionFiscal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tLeyendasFisc_LeyendasFiscales_Leyenda_norma = &tipoXSD{nombre: "LeyendasFisc LeyendasFiscales/Leyenda/@norma", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tLeyendasFisc_LeyendasFiscales_Leyenda_textoLeyenda = &tipoXSD{nombre: "LeyendasFisc LeyendasFiscales/Leyenda/@textoLeyenda", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tIEDU_instEducativas_version = &tipoXSD{nombre: "IEDU instEducativas/@version", prim: "string", ws: "collapse"}
var tIEDU_instEducativas_nombreAlumno = &tipoXSD{nombre: "IEDU instEducativas/@nombreAlumno", prim: "string", ws: "preserve"}
var tIEDU_tCURP = &tipoXSD{nombre: "tCURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron078}},
}}
var tIEDU_instEducativas_nivelEducativo = &tipoXSD{nombre: "IEDU instEducativas/@nivelEducativo", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catIEDU_instEducativas_nivelEducativo},
}}
var tIEDU_instEducativas_autRVOE = &tipoXSD{nombre: "IEDU instEducativas/@autRVOE", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tIEDU_tRFC = &tipoXSD{nombre: "tRFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron079}},
}}
var tDonat11_Donatarias_leyenda = &tipoXSD{nombre: "Donat11 Donatarias/@leyenda", prim: "string", ws: "collapse"}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_IDOrigen = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Origen/@IDOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron044}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NombreRemitente = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Origen/@NombreRemitente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NumRegIdTrib = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Origen/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NombreEstacion = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Origen/@NombreEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NavegacionTrafico = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Origen/@NavegacionTrafico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NavegacionTrafico},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_IDDestino = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Destino/@IDDestino", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron045}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NombreDestinatario = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Destino/@NombreDestinatario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NumRegIdTrib = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Destino/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NombreEstacion = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Destino/@NombreEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NavegacionTrafico = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Destino/@NavegacionTrafico", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NavegacionTrafico},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Ubicaciones/Ubicacion/@DistanciaRecorrida", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron044}},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/CantidadTransporta/@IDDestino", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron045}},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoNeto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@PesoTara", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/DetalleMercancia/@NumPiezas", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCatCartaPorte_cClaveProdSTCC = &tipoXSD{nombre: "c_ClaveProdSTCC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatCartaPorte_cClaveProdSTCC},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_Descripcion = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@Descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 1000},
	{tipo: "pattern", patrones: []patronXSD{patron011}},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_Cantidad = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_Unidad = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron010}},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_Dimensiones = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@Dimensiones", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron080}},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_MaterialPeligroso = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@MaterialPeligroso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte10_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_DescripEmbalaje = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@DescripEmbalaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_PesoEnKg = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@PesoEnKg", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_Mercancia_UUIDComercioExt = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/Mercancia/@UUIDComercioExt", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_IdentificacionVehicular_PlacaVM = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/AutotransporteFederal/IdentificacionVehicular/@PlacaVM", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron081}},
}}
var tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_IdentificacionVehicular_AnioModeloVM = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/AutotransporteFederal/IdentificacionVehicular/@AnioModeloVM", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_Remolques_Remolque_Placa = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/AutotransporteFederal/Remolques/Remolque/@Placa", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron081}},
}}
var tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/AutotransporteFederal/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_NombreAseg = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/AutotransporteFederal/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/AutotransporteFederal/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@MatriculaContenedor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 11},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron058}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/Contenedor/@NumPrecinto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron059}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Matricula = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@Matricula", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 7},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron061}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NumeroOMI", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron062}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@AnioEmbarcacion", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron055}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NombreEmbarc", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@UnidadesDeArqBruto", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumCertITC = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NumCertITC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron068}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Eslora = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@Eslora", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Manga = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@Manga", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Calado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@Calado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@LineaNaviera", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NombreAgenteNaviero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
	{tipo: "pattern", patrones: []patronXSD{patron009}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumViaje = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NumViaje", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteMaritimo/@NumConocEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NumPermisoSCT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@MatriculaAeronave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 5},
	{tipo: "maxLength", n: 10},
	{tipo: "pattern", patrones: []patronXSD{patron069}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NombreAseg = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumeroGuia = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NumeroGuia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron065}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribTranspor = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NumRegIdTribTranspor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NombreTransportista = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NombreTransportista", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NumRegIdTribEmbarc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteAereo/@NombreEmbarcador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/DerechosDePaso/@KilometrajePagado", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoContenedorVacio", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/Carro/Contenedor/@PesoNetoMercancia", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/Carro/@MatriculaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron066}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/Carro/@GuiaCarro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
	{tipo: "pattern", patrones: []patronXSD{patron031}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/Carro/@ToneladasNetasCarro", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/@NombreAseg", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron056}},
}}
var tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/TransporteFerroviario/@NumPolizaSeguro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron057}},
}}
var tCartaPorte10_CartaPorte_Mercancias_PesoBrutoTotal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/@PesoBrutoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_PesoNetoTotal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/@PesoNetoTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
}}
var tCartaPorte10_CartaPorte_Mercancias_NumTotalMercancias = &tipoXSD{nombre: "CartaPorte10 CartaPorte/Mercancias/@NumTotalMercancias", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Operadores/Operador/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Operadores/Operador/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tTdCFDI_tRFCPF = &tipoXSD{nombre: "t_RFC_PF", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron082}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_NumLicencia = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Operadores/Operador/@NumLicencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 16},
	{tipo: "pattern", patrones: []patronXSD{patron067}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_NombreOperador = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Operadores/Operador/@NombreOperador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_NumRegIdTribOperador = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Operadores/Operador/@NumRegIdTribOperador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Propietario/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Propietario/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_NombrePropietario = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Propietario/@NombrePropietario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_NumRegIdTribPropietario = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Propietario/@NumRegIdTribPropietario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Arrendatario/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Arrendatario/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_NombreArrendatario = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Arrendatario/@NombreArrendatario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_NumRegIdTribArrendatario = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Arrendatario/@NumRegIdTribArrendatario", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_Domicilio_Estado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Notificado/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
	{tipo: "pattern", patrones: []patronXSD{patron038}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_Domicilio_CodigoPostal = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Notificado/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_NombreNotificado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Notificado/@NombreNotificado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
	{tipo: "pattern", patrones: []patronXSD{patron017}},
}}
var tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_NumRegIdTribNotificado = &tipoXSD{nombre: "CartaPorte10 CartaPorte/FiguraTransporte/Notificado/@NumRegIdTribNotificado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron040}},
}}
var tCartaPorte10_CartaPorte_TranspInternac = &tipoXSD{nombre: "CartaPorte10 CartaPorte/@TranspInternac", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte10_CartaPorte_TranspInternac},
}}
var tCartaPorte10_CartaPorte_EntradaSalidaMerc = &tipoXSD{nombre: "CartaPorte10 CartaPorte/@EntradaSalidaMerc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCartaPorte10_CartaPorte_EntradaSalidaMerc},
}}
var tCartaPorte10_CartaPorte_TotalDistRec = &tipoXSD{nombre: "CartaPorte10 CartaPorte/@TotalDistRec", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "maxInclusive", valor: "99999"},
}}
var tComExt10_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron083}},
}}
var tComExt10_ComercioExterior_Receptor_NumRegIdTrib = &tipoXSD{nombre: "ComExt10 ComercioExterior/Receptor/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_Calle = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@Calle", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_NumeroExterior = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@NumeroExterior", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 55},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_NumeroInterior = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@NumeroInterior", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 55},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_Colonia = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@Colonia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 120},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_Localidad = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@Localidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 120},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_Referencia = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@Referencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 250},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_Municipio = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@Municipio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 120},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_Estado = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@Estado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
}}
var tComExt10_cPais = &tipoXSD{nombre: "c_Pais", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_cPais},
}}
var tComExt10_ComercioExterior_Destinatario_Domicilio_CodigoPostal = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/Domicilio/@CodigoPostal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 12},
}}
var tComExt10_ComercioExterior_Destinatario_NumRegIdTrib = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
}}
var tComExt10_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tComExt10_ComercioExterior_Destinatario_Nombre = &tipoXSD{nombre: "ComExt10 ComercioExterior/Destinatario/@Nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Marca = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@Marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 35},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Modelo = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@Modelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 80},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_SubModelo = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@SubModelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_NumeroSerie = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/DescripcionesEspecificas/@NumeroSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_NoIdentificacion = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/@NoIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tComExt10_cFraccionArancelaria = &tipoXSD{nombre: "c_FraccionArancelaria", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_cFraccionArancelaria},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_CantidadAduana = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/@CantidadAduana", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "pattern", patrones: []patronXSD{patron071}},
}}
var tComExt10_cUnidadMedidaAduana = &tipoXSD{nombre: "c_UnidadMedidaAduana", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_cUnidadMedidaAduana},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_ValorUnitarioAduana = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/@ValorUnitarioAduana", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "pattern", patrones: []patronXSD{patron074}},
}}
var tComExt10_ComercioExterior_Mercancias_Mercancia_ValorDolares = &tipoXSD{nombre: "ComExt10 ComercioExterior/Mercancias/Mercancia/@ValorDolares", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "pattern", patrones: []patronXSD{patron074}},
}}
var tComExt10_cTipoOperacion = &tipoXSD{nombre: "c_TipoOperacion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_cTipoOperacion},
}}
var tComExt10_cClavePedimento = &tipoXSD{nombre: "c_ClavePedimento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_cClavePedimento},
}}
var tComExt10_ComercioExterior_CertificadoOrigen = &tipoXSD{nombre: "ComExt10 ComercioExterior/@CertificadoOrigen", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_ComercioExterior_CertificadoOrigen},
}}
var tComExt10_ComercioExterior_NumCertificadoOrigen = &tipoXSD{nombre: "ComExt10 ComercioExterior/@NumCertificadoOrigen", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 6},
	{tipo: "maxLength", n: 40},
	{tipo: "pattern", patrones: []patronXSD{patron084}},
}}
var tComExt10_ComercioExterior_NumeroExportadorConfiable = &tipoXSD{nombre: "ComExt10 ComercioExterior/@NumeroExportadorConfiable", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tComExt10_cINCOTERM = &tipoXSD{nombre: "c_INCOTERM", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_cINCOTERM},
}}
var tComExt10_ComercioExterior_Subdivision = &tipoXSD{nombre: "ComExt10 ComercioExterior/@Subdivision", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catComExt10_ComercioExterior_Subdivision},
}}
var tComExt10_ComercioExterior_Observaciones = &tipoXSD{nombre: "ComExt10 ComercioExterior/@Observaciones", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tComExt10_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Impuesto = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@Impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Impuesto},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_TasaoCuota = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@TasaoCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Importe = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "minInclusive", valor: "0.01"},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Identificador = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Identificador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Fecha = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Fecha", prim: "dateTime", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "2015-01-01T00:00:00"},
	{tipo: "pattern", patrones: []patronXSD{patron085}},
}}
var tECC11_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ClaveEstacion = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@ClaveEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tECC11_cTAR = &tipoXSD{nombre: "c_TAR", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catECC11_cTAR},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Cantidad = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
	{tipo: "pattern", patrones: []patronXSD{patron077}},
}}
var tECC11_cclaveProducto = &tipoXSD{nombre: "c_claveProducto", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catECC11_cclaveProducto},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Unidad = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 25},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_NombreCombustible = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@NombreCombustible", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_FolioOperacion = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@FolioOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ValorUnitario = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@ValorUnitario", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "fractionDigits", n: 3},
	{tipo: "pattern", patrones: []patronXSD{patron077}},
}}
var tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Importe = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tECC11_EstadoDeCuentaCombustible_Version = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/@Version", prim: "string", ws: "collapse"}
var tECC11_EstadoDeCuentaCombustible_TipoOperacion = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/@TipoOperacion", prim: "string", ws: "collapse"}
var tECC11_EstadoDeCuentaCombustible_NumeroDeCuenta = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/@NumeroDeCuenta", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tECC11_EstadoDeCuentaCombustible_SubTotal = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/@SubTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tECC11_EstadoDeCuentaCombustible_Total = &tipoXSD{nombre: "ECC11 EstadoDeCuentaCombustible/@Total", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "pattern", patrones: []patronXSD{patron076}},
}}
var tNomina11_Nomina_Percepciones_Percepcion_TipoPercepcion = &tipoXSD{nombre: "Nomina11 Nomina/Percepciones/Percepcion/@TipoPercepcion", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "pattern", patrones: []patronXSD{patron086}},
}}
var tNomina11_Nomina_Percepciones_Percepcion_Clave = &tipoXSD{nombre: "Nomina11 Nomina/Percepciones/Percepcion/@Clave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 15},
}}
var tNomina11_Nomina_Percepciones_Percepcion_Concepto = &tipoXSD{nombre: "Nomina11 Nomina/Percepciones/Percepcion/@Concepto", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNomina11_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tNomina11_Nomina_Deducciones_Deduccion_TipoDeduccion = &tipoXSD{nombre: "Nomina11 Nomina/Deducciones/Deduccion/@TipoDeduccion", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "pattern", patrones: []patronXSD{patron086}},
}}
var tNomina11_Nomina_Deducciones_Deduccion_Clave = &tipoXSD{nombre: "Nomina11 Nomina/Deducciones/Deduccion/@Clave", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 3},
	{tipo: "maxLength", n: 15},
}}
var tNomina11_Nomina_Deducciones_Deduccion_Concepto = &tipoXSD{nombre: "Nomina11 Nomina/Deducciones/Deduccion/@Concepto", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNomina11_Nomina_Incapacidades_Incapacidad_DiasIncapacidad = &tipoXSD{nombre: "Nomina11 Nomina/Incapacidades/Incapacidad/@DiasIncapacidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "1"},
}}
var tNomina11_Nomina_HorasExtras_HorasExtra_TipoHoras = &tipoXSD{nombre: "Nomina11 Nomina/HorasExtras/HorasExtra/@TipoHoras", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catNomina11_Nomina_HorasExtras_HorasExtra_TipoHoras},
}}
var tNomina11_Nomina_RegistroPatronal = &tipoXSD{nombre: "Nomina11 Nomina/@RegistroPatronal", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tNomina11_Nomina_NumEmpleado = &tipoXSD{nombre: "Nomina11 Nomina/@NumEmpleado", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
}}
var tNomina11_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron087}},
}}
var tNomina11_Nomina_TipoRegimen = &tipoXSD{nombre: "Nomina11 Nomina/@TipoRegimen", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tNomina11_Nomina_NumSeguridadSocial = &tipoXSD{nombre: "Nomina11 Nomina/@NumSeguridadSocial", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
}}
var tNomina11_Nomina_NumDiasPagados = &tipoXSD{nombre: "Nomina11 Nomina/@NumDiasPagados", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tNomina11_Nomina_Departamento = &tipoXSD{nombre: "Nomina11 Nomina/@Departamento", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNomina11_tClabe = &tipoXSD{nombre: "t_Clabe", prim: "integer", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron088}},
}}
var tNomina11_Nomina_Banco = &tipoXSD{nombre: "Nomina11 Nomina/@Banco", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "pattern", patrones: []patronXSD{patron086}},
}}
var tNomina11_Nomina_PeriodicidadPago = &tipoXSD{nombre: "Nomina11 Nomina/@PeriodicidadPago", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNomina11_Nomina_RiesgoPuesto = &tipoXSD{nombre: "Nomina11 Nomina/@RiesgoPuesto", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tINE11_INE_Entidad_Contabilidad_IdContabilidad = &tipoXSD{nombre: "INE11 INE/Entidad/Contabilidad/@IdContabilidad", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 6},
}}
var tINE11_tClaveEntidad = &tipoXSD{nombre: "t_ClaveEntidad", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tClaveEntidad},
}}
var tINE11_tAmbito = &tipoXSD{nombre: "t_Ambito", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tAmbito},
}}
var tINE11_INE_Entidad_Ambito = &tipoXSD{nombre: "INE11 INE/Entidad/@Ambito", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tAmbito},
}}
var tINE11_INE_Version = &tipoXSD{nombre: "INE11 INE/@Version", prim: "string", ws: "collapse"}
var tINE11_tTipoProc = &tipoXSD{nombre: "t_TipoProc", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tTipoProc},
}}
var tINE11_INE_TipoProceso = &tipoXSD{nombre: "INE11 INE/@TipoProceso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tTipoProc},
}}
var tINE11_tTipoComite = &tipoXSD{nombre: "t_TipoComite", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tTipoComite},
}}
var tINE11_INE_TipoComite = &tipoXSD{nombre: "INE11 INE/@TipoComite", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE11_tTipoComite},
}}
var tINE11_INE_IdContabilidad = &tipoXSD{nombre: "INE11 INE/@IdContabilidad", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 6},
}}
var tValesDespensa_ValesDeDespensa_Conceptos_Concepto_identificador = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/Conceptos/Concepto/@identificador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tValesDespensa_ValesDeDespensa_Conceptos_Concepto_fecha = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/Conceptos/Concepto/@fecha", prim: "dateTime", ws: "collapse"}
var tValesDespensa_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron079}},
}}
var tValesDespensa_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron078}},
}}
var tValesDespensa_ValesDeDespensa_Conceptos_Concepto_nombre = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/Conceptos/Concepto/@nombre", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tValesDespensa_ValesDeDespensa_Conceptos_Concepto_numSeguridadSocial = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/Conceptos/Concepto/@numSeguridadSocial", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 15},
}}
var tValesDespensa_ValesDeDespensa_Conceptos_Concepto_importe = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/Conceptos/Concepto/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tValesDespensa_ValesDeDespensa_tipoOperacion = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/@tipoOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tValesDespensa_ValesDeDespensa_registroPatronal = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/@registroPatronal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tValesDespensa_ValesDeDespensa_numeroDeCuenta = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/@numeroDeCuenta", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tValesDespensa_ValesDeDespensa_total = &tipoXSD{nombre: "ValesDespensa ValesDeDespensa/@total", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_impuesto = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/Determinados/Determinado/@impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_impuesto},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_tasa = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/Determinados/Determinado/@tasa", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_importe = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/Determinados/Determinado/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_identificador = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@identificador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_fecha = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@fecha", prim: "dateTime", ws: "collapse"}
var tConsumoComb10_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron079}},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_claveEstacion = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@claveEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_cantidad = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@cantidad", prim: "decimal", ws: "collapse"}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_nombreCombustible = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@nombreCombustible", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_folioOperacion = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@folioOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_valorUnitario = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@valorUnitario", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_importe = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb10_ConsumoDeCombustibles_tipoOperacion = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/@tipoOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb10_ConsumoDeCombustibles_numeroDeCuenta = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/@numeroDeCuenta", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb10_ConsumoDeCombustibles_subTotal = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/@subTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb10_ConsumoDeCombustibles_total = &tipoXSD{nombre: "ConsumoComb10 ConsumoDeCombustibles/@total", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_impuesto = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/Determinados/Determinado/@impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_impuesto},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_tasaOCuota = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/Determinados/Determinado/@tasaOCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_importe = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/Determinados/Determinado/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_identificador = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@identificador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_claveEstacion = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@claveEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_cantidad = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@cantidad", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "pattern", patrones: []patronXSD{patron089}},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_nombreCombustible = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@nombreCombustible", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_folioOperacion = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@folioOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_valorUnitario = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@valorUnitario", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_importe = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/Conceptos/ConceptoConsumoDeCombustibles/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb11_ConsumoDeCombustibles_version = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/@version", prim: "string", ws: "collapse"}
var tConsumoComb11_ConsumoDeCombustibles_tipoOperacion = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/@tipoOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb11_ConsumoDeCombustibles_numeroDeCuenta = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/@numeroDeCuenta", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tConsumoComb11_ConsumoDeCombustibles_subTotal = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/@subTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tConsumoComb11_ConsumoDeCombustibles_total = &tipoXSD{nombre: "ConsumoComb11 ConsumoDeCombustibles/@total", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tPagoEnEspecie_PagoEnEspecie_CvePIC = &tipoXSD{nombre: "PagoEnEspecie PagoEnEspecie/@CvePIC", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 25},
	{tipo: "pattern", patrones: []patronXSD{patron090}},
}}
var tPagoEnEspecie_PagoEnEspecie_FolioSolDon = &tipoXSD{nombre: "PagoEnEspecie PagoEnEspecie/@FolioSolDon", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 11},
	{tipo: "pattern", patrones: []patronXSD{patron091}},
}}
var tPagoEnEspecie_PagoEnEspecie_PzaArtNombre = &tipoXSD{nombre: "PagoEnEspecie PagoEnEspecie/@PzaArtNombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tPagoEnEspecie_PagoEnEspecie_PzaArtTecn = &tipoXSD{nombre: "PagoEnEspecie PagoEnEspecie/@PzaArtTecn", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tPagoEnEspecie_PagoEnEspecie_PzaArtAProd = &tipoXSD{nombre: "PagoEnEspecie PagoEnEspecie/@PzaArtAProd", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron092}},
}}
var tPagoEnEspecie_PagoEnEspecie_PzaArtDim = &tipoXSD{nombre: "PagoEnEspecie PagoEnEspecie/@PzaArtDim", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_TipoInmueble = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@TipoInmueble", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron093}},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Calle = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@Calle", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_NoExterior = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@NoExterior", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 55},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_NoInterior = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@NoInterior", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Colonia = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@Colonia", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Localidad = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@Localidad", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Referencia = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@Referencia", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Municipio = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@Municipio", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tNotarios_tEntidadFederativa = &tipoXSD{nombre: "t_EntidadFederativa", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catNotarios_tEntidadFederativa},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Pais = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@Pais", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron094}},
}}
var tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_CodigoPostal = &tipoXSD{nombre: "Notarios NotariosPublicos/DescInmuebles/DescInmueble/@CodigoPostal", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron005}},
}}
var tNotarios_NotariosPublicos_DatosOperacion_NumInstrumentoNotarial = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosOperacion/@NumInstrumentoNotarial", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "maxInclusive", valor: "999999"},
}}
var tNotarios_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tNotarios_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron078}},
}}
var tNotarios_NotariosPublicos_DatosNotario_NumNotaria = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosNotario/@NumNotaria", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "maxInclusive", valor: "999"},
}}
var tNotarios_NotariosPublicos_DatosNotario_Adscripcion = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosNotario/@Adscripcion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 255},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante_Nombre = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosUnEnajenante/@Nombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante_ApellidoPaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosUnEnajenante/@ApellidoPaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante_ApellidoMaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosUnEnajenante/@ApellidoMaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron079}},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_Nombre = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosEnajenantesCopSC/DatosEnajenanteCopSC/@Nombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_ApellidoPaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosEnajenantesCopSC/DatosEnajenanteCopSC/@ApellidoPaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_ApellidoMaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosEnajenantesCopSC/DatosEnajenanteCopSC/@ApellidoMaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_Porcentaje = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/DatosEnajenantesCopSC/DatosEnajenanteCopSC/@Porcentaje", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "maxInclusive", valor: "100.00"},
	{tipo: "fractionDigits", n: 2},
}}
var tNotarios_NotariosPublicos_DatosEnajenante_CoproSocConyugalE = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosEnajenante/@CoproSocConyugalE", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catNotarios_NotariosPublicos_DatosEnajenante_CoproSocConyugalE},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente_Nombre = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosUnAdquiriente/@Nombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente_ApellidoPaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosUnAdquiriente/@ApellidoPaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente_ApellidoMaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosUnAdquiriente/@ApellidoMaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_Nombre = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosAdquirientesCopSC/DatosAdquirienteCopSC/@Nombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_ApellidoPaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosAdquirientesCopSC/DatosAdquirienteCopSC/@ApellidoPaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_ApellidoMaterno = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosAdquirientesCopSC/DatosAdquirienteCopSC/@ApellidoMaterno", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 200},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_Porcentaje = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/DatosAdquirientesCopSC/DatosAdquirienteCopSC/@Porcentaje", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.00"},
	{tipo: "maxInclusive", valor: "100.00"},
	{tipo: "fractionDigits", n: 2},
}}
var tNotarios_NotariosPublicos_DatosAdquiriente_CoproSocConyugalE = &tipoXSD{nombre: "Notarios NotariosPublicos/DatosAdquiriente/@CoproSocConyugalE", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catNotarios_NotariosPublicos_DatosAdquiriente_CoproSocConyugalE},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_Calle = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@Calle", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 150},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_NoExterior = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@NoExterior", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 55},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_NoInterior = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@NoInterior", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 30},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_Colonia = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@Colonia", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_Localidad = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@Localidad", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_Referencia = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@Referencia", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_Municipio = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@Municipio", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tParcialesConstruccion_tEntidadFederativa = &tipoXSD{nombre: "t_EntidadFederativa", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catParcialesConstruccion_tEntidadFederativa},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_Estado = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@Estado", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catParcialesConstruccion_tEntidadFederativa},
	{tipo: "length", n: 2},
}}
var tParcialesConstruccion_parcialesconstruccion_Inmueble_CodigoPostal = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/Inmueble/@CodigoPostal", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron005}},
}}
var tParcialesConstruccion_parcialesconstruccion_Version = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/@Version", prim: "string", ws: "collapse"}
var tParcialesConstruccion_parcialesconstruccion_NumPerLicoAut = &tipoXSD{nombre: "ParcialesConstruccion parcialesconstruccion/@NumPerLicoAut", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tVehiculoUsado_VehiculoUsado_InformacionAduanera_numero = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVehiculoUsado_VehiculoUsado_InformacionAduanera_fecha = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tVehiculoUsado_VehiculoUsado_InformacionAduanera_aduana = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVehiculoUsado_VehiculoUsado_Version = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@Version", prim: "string", ws: "collapse"}
var tVehiculoUsado_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tVehiculoUsado_VehiculoUsado_claveVehicular = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@claveVehicular", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 7},
}}
var tVehiculoUsado_VehiculoUsado_marca = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tVehiculoUsado_VehiculoUsado_tipo = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@tipo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tVehiculoUsado_VehiculoUsado_modelo = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@modelo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron092}},
}}
var tVehiculoUsado_VehiculoUsado_numeroMotor = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@numeroMotor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tVehiculoUsado_VehiculoUsado_numeroSerie = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@numeroSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tVehiculoUsado_VehiculoUsado_NIV = &tipoXSD{nombre: "VehiculoUsado VehiculoUsado/@NIV", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_Via = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/datosTransito/@Via", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_Via},
}}
var tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_TipoId = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/datosTransito/@TipoId", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_NumeroId = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/datosTransito/@NumeroId", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_Nacionalidad = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/datosTransito/@Nacionalidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_EmpresaTransporte = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/datosTransito/@EmpresaTransporte", prim: "string", ws: "collapse"}
var tTuristaPasajero_TuristaPasajeroExtranjero_fechadeTransito = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/@fechadeTransito", prim: "dateTime", ws: "collapse"}
var tTuristaPasajero_TuristaPasajeroExtranjero_tipoTransito = &tipoXSD{nombre: "TuristaPasajero TuristaPasajeroExtranjero/@tipoTransito", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catTuristaPasajero_TuristaPasajeroExtranjero_tipoTransito},
}}
var tAcreditamientoIEPS_cTAR = &tipoXSD{nombre: "c_TAR", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catAcreditamientoIEPS_cTAR},
}}
var tAcreditamientoIEPS_acreditamientoIEPS_TAR = &tipoXSD{nombre: "AcreditamientoIEPS acreditamientoIEPS/@TAR", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catAcreditamientoIEPS_cTAR},
	{tipo: "length", n: 3},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Ordenante_BancoEmisor = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Ordenante/@BancoEmisor", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Ordenante_Nombre = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Ordenante/@Nombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "maxLength", n: 40},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Ordenante_TipoCuenta = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Ordenante/@TipoCuenta", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 2},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Ordenante_Cuenta = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Ordenante/@Cuenta", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 20},
}}
var tSPEI_tRFC = &tipoXSD{nombre: "tRFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 2},
	{tipo: "maxLength", n: 18},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_BancoReceptor = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@BancoReceptor", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_Nombre = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@Nombre", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "maxLength", n: 40},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_TipoCuenta = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@TipoCuenta", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 2},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_Cuenta = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@Cuenta", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 20},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_Concepto = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@Concepto", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "maxLength", n: 40},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_IVA = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@IVA", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
	{tipo: "totalDigits", n: 19},
}}
var tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_MontoPago = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/Beneficiario/@MontoPago", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "totalDigits", n: 19},
	{tipo: "fractionDigits", n: 2},
}}
var tSPEI_ComplementoSPEI_SPEITercero_FechaOperacion = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/@FechaOperacion", prim: "date", ws: "collapse", facetas: []facetaXSD{
	{tipo: "maxInclusive", valor: "9999-12-31"},
	{tipo: "minInclusive", valor: "0001-01-01"},
	{tipo: "pattern", patrones: []patronXSD{patron095}},
}}
var xstime = &tipoXSD{nombre: "xs:time", prim: "time", ws: "collapse"}
var tSPEI_ComplementoSPEI_SPEITercero_Hora = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/@Hora", prim: "time", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron096}},
}}
var tSPEI_ComplementoSPEI_SPEITercero_ClaveSPEI = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/@ClaveSPEI", prim: "integer", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 5},
}}
var tSPEI_ComplementoSPEI_SPEITercero_sello = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/@sello", prim: "string", ws: "collapse"}
var tSPEI_ComplementoSPEI_SPEITercero_numeroCertificado = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/@numeroCertificado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "maxLength", n: 20},
}}
var tSPEI_ComplementoSPEI_SPEITercero_cadenaCDA = &tipoXSD{nombre: "SPEI ComplementoSPEI/SPEITercero/@cadenaCDA", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "maxLength", n: 841},
}}
var tRenovacion_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRenovacion_cTipoVehiculor = &tipoXSD{nombre: "c_TipoVehiculo_r", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRenovacion_cTipoVehiculor},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Marca = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@Marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_TipooClase = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@TipooClase", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Anio = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@Año", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1900"},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Modelo = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@Modelo", prim: "string", ws: "collapse"}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NIV = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@NIV", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumSerie = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@NumSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumPlacas = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@NumPlacas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumMotor = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@NumMotor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumFolTarjCir = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@NumFolTarjCir", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumPedIm = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@NumPedIm", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Aduana = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@Aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Foliofiscal = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculosUsadosEnajenadoPermAlFab/@Foliofiscal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 36},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Anio = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculoNuvoSemEnajenadoFabAlPerm/@Año", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1900"},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Modelo = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculoNuvoSemEnajenadoFabAlPerm/@Modelo", prim: "string", ws: "collapse"}
var tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_NumPlacas = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoRenovVehicular/VehiculoNuvoSemEnajenadoFabAlPerm/@NumPlacas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tRenovacion_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron097}},
}}
var tRenovacion_cVehiculoEnajenado = &tipoXSD{nombre: "c_VehiculoEnajenado", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRenovacion_cVehiculoEnajenado},
}}
var tRenovacion_cTipoVehiculos = &tipoXSD{nombre: "c_TipoVehiculo_s", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRenovacion_cTipoVehiculos},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Marca = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@Marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_TipooClase = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@TipooClase", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Anio = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@Año", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1900"},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Modelo = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@Modelo", prim: "string", ws: "collapse"}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NIV = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NIV", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumSerie = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NumSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumPlacas = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NumPlacas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumMotor = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NumMotor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumFolTarjCir = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NumFolTarjCir", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumFolAvisoint = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NumFolAvisoint", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumPedIm = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@NumPedIm", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Aduana = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@Aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Foliofiscal = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoUsadoEnajenadoPermAlFab/@Foliofiscal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 36},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Anio = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoNuvoSemEnajenadoFabAlPerm/@Año", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1900"},
}}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Modelo = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoNuvoSemEnajenadoFabAlPerm/@Modelo", prim: "string", ws: "collapse"}
var tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_NumPlacas = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/DecretoSustitVehicular/VehiculoNuvoSemEnajenadoFabAlPerm/@NumPlacas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tRenovacion_renovacionysustitucionvehiculos_Version = &tipoXSD{nombre: "Renovacion renovacionysustitucionvehiculos/@Version", prim: "string", ws: "collapse"}
var tRenovacion_cTipoDecreto = &tipoXSD{nombre: "c_TipoDecreto", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRenovacion_cTipoDecreto},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_Marca = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@Marca", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_TipooClase = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@TipooClase", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_Anio = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@Año", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1900"},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_Modelo = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@Modelo", prim: "string", ws: "collapse"}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NIV = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@NIV", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumSerie = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@NumSerie", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumPlacas = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@NumPlacas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumMotor = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@NumMotor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 17},
}}
var tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumFolTarjCir = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/VehiculoDestruido/@NumFolTarjCir", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tCertDestruccion_certificadodedestruccion_InformacionAduanera_NumPedImp = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/InformacionAduanera/@NumPedImp", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 40},
}}
var tCertDestruccion_certificadodedestruccion_InformacionAduanera_Aduana = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/InformacionAduanera/@Aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tCertDestruccion_certificadodedestruccion_Version = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/@Version", prim: "string", ws: "collapse"}
var tCertDestruccion_cTipoSerie = &tipoXSD{nombre: "c_TipoSerie", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCertDestruccion_cTipoSerie},
}}
var tCertDestruccion_certificadodedestruccion_NumFolDesVeh = &tipoXSD{nombre: "CertDestruccion certificadodedestruccion/@NumFolDesVeh", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tObrasArte_cTipoBien = &tipoXSD{nombre: "c_TipoBien", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catObrasArte_cTipoBien},
}}
var tObrasArte_obrasarteantiguedades_TipoBien = &tipoXSD{nombre: "ObrasArte obrasarteantiguedades/@TipoBien", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catObrasArte_cTipoBien},
	{tipo: "length", n: 2},
}}
var tObrasArte_obrasarteantiguedades_OtrosTipoBien = &tipoXSD{nombre: "ObrasArte obrasarteantiguedades/@OtrosTipoBien", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 255},
}}
var tObrasArte_cTituloadquirido = &tipoXSD{nombre: "c_Tituloadquirido", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catObrasArte_cTituloadquirido},
}}
var tObrasArte_obrasarteantiguedades_TituloAdquirido = &tipoXSD{nombre: "ObrasArte obrasarteantiguedades/@TituloAdquirido", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catObrasArte_cTituloadquirido},
	{tipo: "length", n: 2},
}}
var tObrasArte_obrasarteantiguedades_OtrosTituloAdquirido = &tipoXSD{nombre: "ObrasArte obrasarteantiguedades/@OtrosTituloAdquirido", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 255},
}}
var tObrasArte_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tObrasArte_obrasarteantiguedades_FechaAdquisicion = &tipoXSD{nombre: "ObrasArte obrasarteantiguedades/@FechaAdquisicion", prim: "date", ws: "collapse"}
var tObrasArte_cCaracteristicasDeObraoPieza = &tipoXSD{nombre: "c_CaracterísticasDeObraoPieza", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catObrasArte_cCaracteristicasDeObraoPieza},
}}
var tObrasArte_obrasarteantiguedades_CaracteristicasDeObraoPieza = &tipoXSD{nombre: "ObrasArte obrasarteantiguedades/@CaracterísticasDeObraoPieza", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catObrasArte_cCaracteristicasDeObraoPieza},
	{tipo: "length", n: 2},
}}
var tPFIC_PFintegranteCoordinado_ClaveVehicular = &tipoXSD{nombre: "PFIC PFintegranteCoordinado/@ClaveVehicular", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tPFIC_PFintegranteCoordinado_Placa = &tipoXSD{nombre: "PFIC PFintegranteCoordinado/@Placa", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tPFIC_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron079}},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_OrigenErogacion = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/DocumentoRelacionado/@OrigenErogacion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_OrigenErogacion},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_FolioFiscalVinculado = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/DocumentoRelacionado/@FolioFiscalVinculado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_RFCProveedor = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/DocumentoRelacionado/@RFCProveedor", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_NumeroPedimentoVinculado = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/DocumentoRelacionado/@NumeroPedimentoVinculado", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "length", n: 21},
	{tipo: "pattern", patrones: []patronXSD{patron007}},
}}
var tCatHidrocarburos_ClavePedimento = &tipoXSD{nombre: "ClavePedimento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidrocarburos_ClavePedimento},
}}
var tCatHidrocarburos_ClavePagoPedimento = &tipoXSD{nombre: "ClavePagoPedimento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidrocarburos_ClavePagoPedimento},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_FechaFolioFiscalVinculado = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/DocumentoRelacionado/@FechaFolioFiscalVinculado", prim: "date", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron098}},
}}
var tCatHidrocarburos_Meses = &tipoXSD{nombre: "Meses", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidrocarburos_Meses},
}}
var tCatHidrocarburos_Tareas = &tipoXSD{nombre: "Tareas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidrocarburos_Tareas},
}}
var tCatHidrocarburos_SubActividad = &tipoXSD{nombre: "SubActividad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidrocarburos_SubActividad},
}}
var tCatHidrocarburos_Actividades = &tipoXSD{nombre: "Actividades", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatHidrocarburos_Actividades},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos_Pozos_Pozo = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/CentroCostos/Yacimientos/Pozos/@Pozo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos_Yacimiento = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/CentroCostos/Yacimientos/@Yacimiento", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Campo = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/CentroCostos/@Campo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_TipoErogacion = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/@TipoErogacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catGastosHidro_GastosHidrocarburos_Erogacion_TipoErogacion},
}}
var tGastosHidro_GastosHidrocarburos_Erogacion_Porcentaje = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/Erogacion/@Porcentaje", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "maxInclusive", valor: "99.999"},
	{tipo: "minInclusive", valor: "00.001"},
}}
var tGastosHidro_GastosHidrocarburos_Version = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/@Version", prim: "string", ws: "collapse"}
var tGastosHidro_GastosHidrocarburos_NumeroContrato = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/@NumeroContrato", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tGastosHidro_GastosHidrocarburos_AreaContractual = &tipoXSD{nombre: "GastosHidro GastosHidrocarburos/@AreaContractual", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tIngresosHidro_IngresosHidrocarburos_DocumentoRelacionado_FolioFiscalVinculado = &tipoXSD{nombre: "IngresosHidro IngresosHidrocarburos/DocumentoRelacionado/@FolioFiscalVinculado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tIngresosHidro_IngresosHidrocarburos_DocumentoRelacionado_FechaFolioFiscalVinculado = &tipoXSD{nombre: "IngresosHidro IngresosHidrocarburos/DocumentoRelacionado/@FechaFolioFiscalVinculado", prim: "date", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron098}},
}}
var tIngresosHidro_IngresosHidrocarburos_Version = &tipoXSD{nombre: "IngresosHidro IngresosHidrocarburos/@Version", prim: "string", ws: "collapse"}
var tIngresosHidro_IngresosHidrocarburos_NumeroContrato = &tipoXSD{nombre: "IngresosHidro IngresosHidrocarburos/@NumeroContrato", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 50},
	{tipo: "pattern", patrones: []patronXSD{patron041}},
}}
var tIngresosHidro_IngresosHidrocarburos_Porcentaje = &tipoXSD{nombre: "IngresosHidro IngresosHidrocarburos/@Porcentaje", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 3},
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "maxInclusive", valor: "100.000"},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_fecha = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@fecha", prim: "dateTime", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_referencia = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@referencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_descripcion = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_importe = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_moneda = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@moneda", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_saldoInicial = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@saldoInicial", prim: "decimal", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_saldoAlCorte = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECB/@saldoAlCorte", prim: "decimal", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_fecha = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@fecha", prim: "dateTime", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_referencia = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@referencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_descripcion = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_RFCenajenante = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@RFCenajenante", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron099}},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_Importe = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_moneda = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@moneda", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_saldoInicial = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@saldoInicial", prim: "decimal", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_saldoAlCorte = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/Movimientos/MovimientoECBFiscal/@saldoAlCorte", prim: "decimal", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_version = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/@version", prim: "string", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_numeroCuenta = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/@numeroCuenta", prim: "integer", ws: "collapse"}
var tECB_EstadoDeCuentaBancario_nombreCliente = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/@nombreCliente", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_periodo = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/@periodo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECB_EstadoDeCuentaBancario_sucursal = &tipoXSD{nombre: "ECB EstadoDeCuentaBancario/@sucursal", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_impuesto = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_impuesto},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_tasa = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@tasa", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_importe = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/Traslados/Traslado/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_identificador = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@identificador", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_fecha = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@fecha", prim: "dateTime", ws: "collapse"}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_rfc = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@rfc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_claveEstacion = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@claveEstacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 10},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_cantidad = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@cantidad", prim: "decimal", ws: "collapse"}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_nombreCombustible = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@nombreCombustible", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_folioOperacion = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@folioOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_valorUnitario = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@valorUnitario", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_importe = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible/@importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECC10_EstadoDeCuentaCombustible_tipoOperacion = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/@tipoOperacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC10_EstadoDeCuentaCombustible_numeroDeCuenta = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/@numeroDeCuenta", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tECC10_EstadoDeCuentaCombustible_subTotal = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/@subTotal", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tECC10_EstadoDeCuentaCombustible_total = &tipoXSD{nombre: "ECC10 EstadoDeCuentaCombustible/@total", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 2},
}}
var tRet20_Retenciones_CfdiRetenRelacionados_UUID = &tipoXSD{nombre: "Ret20 Retenciones/CfdiRetenRelacionados/@UUID", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tRet20_Retenciones_Emisor_NomDenRazSocE = &tipoXSD{nombre: "Ret20 Retenciones/Emisor/@NomDenRazSocE", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
}}
var tRet20_Retenciones_Receptor_Nacional_NomDenRazSocR = &tipoXSD{nombre: "Ret20 Retenciones/Receptor/Nacional/@NomDenRazSocR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 254},
}}
var tRet20_Retenciones_Receptor_Nacional_DomicilioFiscalR = &tipoXSD{nombre: "Ret20 Retenciones/Receptor/Nacional/@DomicilioFiscalR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 5},
	{tipo: "pattern", patrones: []patronXSD{patron005}},
}}
var tRet20_Retenciones_Receptor_Extranjero_NumRegIdTribR = &tipoXSD{nombre: "Ret20 Retenciones/Receptor/Extranjero/@NumRegIdTribR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tRet20_Retenciones_Receptor_Extranjero_NomDenRazSocR = &tipoXSD{nombre: "Ret20 Retenciones/Receptor/Extranjero/@NomDenRazSocR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRet20_Retenciones_Receptor_NacionalidadR = &tipoXSD{nombre: "Ret20 Retenciones/Receptor/@NacionalidadR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRet20_Retenciones_Receptor_NacionalidadR},
}}
var tCatRetenciones_cPeriodo = &tipoXSD{nombre: "c_Periodo", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cPeriodo},
}}
var tCatRetenciones_cEjercicio = &tipoXSD{nombre: "c_Ejercicio", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cEjercicio},
}}
var tCatRetenciones_cTipoPagoRet = &tipoXSD{nombre: "c_TipoPagoRet", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cTipoPagoRet},
}}
var tRet20_Retenciones_Version = &tipoXSD{nombre: "Ret20 Retenciones/@Version", prim: "string", ws: "collapse"}
var tRet20_Retenciones_FolioInt = &tipoXSD{nombre: "Ret20 Retenciones/@FolioInt", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron100}},
}}
var tRet20_Retenciones_Sello = &tipoXSD{nombre: "Ret20 Retenciones/@Sello", prim: "string", ws: "collapse"}
var tRet20_Retenciones_NoCertificado = &tipoXSD{nombre: "Ret20 Retenciones/@NoCertificado", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 20},
	{tipo: "pattern", patrones: []patronXSD{patron015}},
}}
var tRet20_Retenciones_Certificado = &tipoXSD{nombre: "Ret20 Retenciones/@Certificado", prim: "string", ws: "collapse"}
var tCatRetenciones_cCveRetenc = &tipoXSD{nombre: "c_CveRetenc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cCveRetenc},
}}
var tRet20_Retenciones_DescRetenc = &tipoXSD{nombre: "Ret20 Retenciones/@DescRetenc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tRet10_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron097}},
}}
var tRet10_Retenciones_Emisor_NomDenRazSocE = &tipoXSD{nombre: "Ret10 Retenciones/Emisor/@NomDenRazSocE", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRet10_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron087}},
}}
var tRet10_Retenciones_Receptor_Nacional_RFCRecep = &tipoXSD{nombre: "Ret10 Retenciones/Receptor/Nacional/@RFCRecep", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron097}},
}}
var tRet10_Retenciones_Receptor_Nacional_NomDenRazSocR = &tipoXSD{nombre: "Ret10 Retenciones/Receptor/Nacional/@NomDenRazSocR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRet10_Retenciones_Receptor_Extranjero_NumRegIdTrib = &tipoXSD{nombre: "Ret10 Retenciones/Receptor/Extranjero/@NumRegIdTrib", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tRet10_Retenciones_Receptor_Extranjero_NomDenRazSocR = &tipoXSD{nombre: "Ret10 Retenciones/Receptor/Extranjero/@NomDenRazSocR", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRet10_Retenciones_Receptor_Nacionalidad = &tipoXSD{nombre: "Ret10 Retenciones/Receptor/@Nacionalidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRet10_Retenciones_Receptor_Nacionalidad},
}}
var tRet10_Retenciones_Periodo_MesIni = &tipoXSD{nombre: "Ret10 Retenciones/Periodo/@MesIni", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "maxInclusive", valor: "12"},
}}
var tRet10_Retenciones_Periodo_MesFin = &tipoXSD{nombre: "Ret10 Retenciones/Periodo/@MesFin", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
	{tipo: "maxInclusive", valor: "12"},
}}
var tRet10_Retenciones_Periodo_Ejerc = &tipoXSD{nombre: "Ret10 Retenciones/Periodo/@Ejerc", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "2004"},
	{tipo: "maxInclusive", valor: "2024"},
}}
var tRet10_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tCatRetenciones_cTipoImpuesto = &tipoXSD{nombre: "c_TipoImpuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cTipoImpuesto},
}}
var tRet10_Retenciones_Totales_ImpRetenidos_Impuesto = &tipoXSD{nombre: "Ret10 Retenciones/Totales/ImpRetenidos/@Impuesto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cTipoImpuesto},
}}
var tRet10_Retenciones_Totales_ImpRetenidos_TipoPagoRet = &tipoXSD{nombre: "Ret10 Retenciones/Totales/ImpRetenidos/@TipoPagoRet", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRet10_Retenciones_Totales_ImpRetenidos_TipoPagoRet},
}}
var tRet10_Retenciones_Version = &tipoXSD{nombre: "Ret10 Retenciones/@Version", prim: "string", ws: "collapse"}
var tRet10_Retenciones_FolioInt = &tipoXSD{nombre: "Ret10 Retenciones/@FolioInt", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tRet10_Retenciones_Sello = &tipoXSD{nombre: "Ret10 Retenciones/@Sello", prim: "string", ws: "collapse"}
var tRet10_Retenciones_NumCert = &tipoXSD{nombre: "Ret10 Retenciones/@NumCert", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 20},
}}
var tRet10_Retenciones_Cert = &tipoXSD{nombre: "Ret10 Retenciones/@Cert", prim: "string", ws: "collapse"}
var tRet10_Retenciones_FechaExp = &tipoXSD{nombre: "Ret10 Retenciones/@FechaExp", prim: "dateTime", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "2014-01-01T00:00:00-06:00"},
	{tipo: "pattern", patrones: []patronXSD{patron101}},
}}
var tCatRetenciones_cRetenciones = &tipoXSD{nombre: "c_Retenciones", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cRetenciones},
}}
var tRet10_Retenciones_DescRetenc = &tipoXSD{nombre: "Ret10 Retenciones/@DescRetenc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ImpuestosTrasladadosdelServicio_Base = &tipoXSD{nombre: "RetPlataformas10 ServiciosPlataformasTecnologicas/Servicios/DetallesDelServicio/ImpuestosTrasladadosdelServicio/@Base", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
	{tipo: "pattern", patrones: []patronXSD{patron006}},
}}
var tCatPlataformas_cTasaCuota = &tipoXSD{nombre: "c_TasaCuota", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatPlataformas_cTasaCuota},
}}
var tCatRetenciones_cEntidadesFederativas = &tipoXSD{nombre: "c_EntidadesFederativas", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cEntidadesFederativas},
}}
var tRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ComisionDelServicio_Base = &tipoXSD{nombre: "RetPlataformas10 ServiciosPlataformasTecnologicas/Servicios/DetallesDelServicio/ComisionDelServicio/@Base", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "minInclusive", valor: "0.000001"},
	{tipo: "pattern", patrones: []patronXSD{patron006}},
}}
var tRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ComisionDelServicio_Porcentaje = &tipoXSD{nombre: "RetPlataformas10 ServiciosPlataformasTecnologicas/Servicios/DetallesDelServicio/ComisionDelServicio/@Porcentaje", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "0.001"},
	{tipo: "maxInclusive", valor: "1.0"},
	{tipo: "fractionDigits", n: 3},
}}
var tCatPlataformas_cFormaPagoServ = &tipoXSD{nombre: "c_FormaPagoServ", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatPlataformas_cFormaPagoServ},
}}
var tCatPlataformas_cTipoDeServ = &tipoXSD{nombre: "c_TipoDeServ", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatPlataformas_cTipoDeServ},
}}
var tCatPlataformas_cSubTipoServ = &tipoXSD{nombre: "c_SubTipoServ", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatPlataformas_cSubTipoServ},
}}
var tCatRetenciones_cPeriodicidad = &tipoXSD{nombre: "c_Periodicidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cPeriodicidad},
}}
var tRetPlataformas10_ServiciosPlataformasTecnologicas_NumServ = &tipoXSD{nombre: "RetPlataformas10 ServiciosPlataformasTecnologicas/@NumServ", prim: "integer", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minInclusive", valor: "1"},
}}
var tRetIntereses_Intereses_SistFinanciero = &tipoXSD{nombre: "RetIntereses Intereses/@SistFinanciero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetIntereses_Intereses_SistFinanciero},
}}
var tRetIntereses_Intereses_RetiroAORESRetInt = &tipoXSD{nombre: "RetIntereses Intereses/@RetiroAORESRetInt", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetIntereses_Intereses_RetiroAORESRetInt},
}}
var tRetIntereses_Intereses_OperFinancDerivad = &tipoXSD{nombre: "RetIntereses Intereses/@OperFinancDerivad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetIntereses_Intereses_OperFinancDerivad},
}}
var tRetIntereses_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tCatRetenciones_cPais = &tipoXSD{nombre: "c_Pais", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cPais},
}}
var tCatRetenciones_cTipoContribuyenteSujetoRetencion = &tipoXSD{nombre: "c_TipoContribuyenteSujetoRetencion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cTipoContribuyenteSujetoRetencion},
}}
var tRetPagosExtranjeros_Pagosaextranjeros_NoBeneficiario_DescripcionConcepto = &tipoXSD{nombre: "RetPagosExtranjeros Pagosaextranjeros/NoBeneficiario/@DescripcionConcepto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 255},
}}
var tRetPagosExtranjeros_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron097}},
}}
var tRetPagosExtranjeros_tCURP = &tipoXSD{nombre: "t_CURP", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 18},
	{tipo: "pattern", patrones: []patronXSD{patron078}},
}}
var tRetPagosExtranjeros_Pagosaextranjeros_Beneficiario_NomDenRazSocB = &tipoXSD{nombre: "RetPagosExtranjeros Pagosaextranjeros/Beneficiario/@NomDenRazSocB", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRetPagosExtranjeros_Pagosaextranjeros_Beneficiario_DescripcionConcepto = &tipoXSD{nombre: "RetPagosExtranjeros Pagosaextranjeros/Beneficiario/@DescripcionConcepto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 255},
}}
var tRetPagosExtranjeros_Pagosaextranjeros_EsBenefEfectDelCobro = &tipoXSD{nombre: "RetPagosExtranjeros Pagosaextranjeros/@EsBenefEfectDelCobro", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPagosExtranjeros_Pagosaextranjeros_EsBenefEfectDelCobro},
}}
var tCatRetenciones_cTipoDividendoOUtilidadDistribuida = &tipoXSD{nombre: "c_TipoDividendoOUtilidadDistribuida", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatRetenciones_cTipoDividendoOUtilidadDistribuida},
}}
var tRetDividendos_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetDividendos_Dividendos_DividOUtil_TipoSocDistrDiv = &tipoXSD{nombre: "RetDividendos Dividendos/DividOUtil/@TipoSocDistrDiv", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetDividendos_Dividendos_DividOUtil_TipoSocDistrDiv},
}}
var tRetDividendos_Dividendos_Remanente_ProporcionRem = &tipoXSD{nombre: "RetDividendos Dividendos/Remanente/@ProporcionRem", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
	{tipo: "maxInclusive", valor: "9999999999"},
}}
var tRetEnajenacion_EnajenaciondeAcciones_ContratoIntermediacion = &tipoXSD{nombre: "RetEnajenacion EnajenaciondeAcciones/@ContratoIntermediacion", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRetEnajenacion_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetArrendamiento_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetDerivados_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetDerivados_Operacionesconderivados_MontGanAcum = &tipoXSD{nombre: "RetDerivados Operacionesconderivados/@MontGanAcum", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetHipotecarios_Intereseshipotecarios_CreditoDeInstFinanc = &tipoXSD{nombre: "RetHipotecarios Intereseshipotecarios/@CreditoDeInstFinanc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetHipotecarios_Intereseshipotecarios_CreditoDeInstFinanc},
}}
var tRetHipotecarios_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetHipotecarios_Intereseshipotecarios_NumContrato = &tipoXSD{nombre: "RetHipotecarios Intereseshipotecarios/@NumContrato", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas_IntegracIngresos_Concepto = &tipoXSD{nombre: "RetFideicomiso Fideicomisonoempresarial/IngresosOEntradas/IntegracIngresos/@Concepto", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tRetFideicomiso_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas_PropDelMontTot = &tipoXSD{nombre: "RetFideicomiso Fideicomisonoempresarial/IngresosOEntradas/@PropDelMontTot", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "maxInclusive", valor: "9999999999"},
	{tipo: "fractionDigits", n: 6},
}}
var tRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas_IntegracEgresos_ConceptoS = &tipoXSD{nombre: "RetFideicomiso Fideicomisonoempresarial/DeduccOSalidas/IntegracEgresos/@ConceptoS", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas_PropDelMontTot = &tipoXSD{nombre: "RetFideicomiso Fideicomisonoempresarial/DeduccOSalidas/@PropDelMontTot", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "pattern", patrones: []patronXSD{patron102}},
}}
var tRetFideicomiso_Fideicomisonoempresarial_RetEfectFideicomiso_DescRetRelPagFideic = &tipoXSD{nombre: "RetFideicomiso Fideicomisonoempresarial/RetEfectFideicomiso/@DescRetRelPagFideic", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tRetPlanesRetiro10_Planesderetiro_SistemaFinanc = &tipoXSD{nombre: "RetPlanesRetiro10 Planesderetiro/@SistemaFinanc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPlanesRetiro10_Planesderetiro_SistemaFinanc},
}}
var tRetPlanesRetiro10_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetPlanesRetiro10_Planesderetiro_HuboRetirosAnioInmAntPer = &tipoXSD{nombre: "RetPlanesRetiro10 Planesderetiro/@HuboRetirosAnioInmAntPer", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPlanesRetiro10_Planesderetiro_HuboRetirosAnioInmAntPer},
}}
var tRetPlanesRetiro10_Planesderetiro_HuboRetirosAnioInmAnt = &tipoXSD{nombre: "RetPlanesRetiro10 Planesderetiro/@HuboRetirosAnioInmAnt", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPlanesRetiro10_Planesderetiro_HuboRetirosAnioInmAnt},
}}
var tCatPlanesRetiro_cTipoAportODep = &tipoXSD{nombre: "c_TipoAportODep", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catCatPlanesRetiro_cTipoAportODep},
}}
var tRetPlanesRetiro11_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetPlanesRetiro11_tRFC = &tipoXSD{nombre: "t_RFC", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 12},
	{tipo: "maxLength", n: 13},
	{tipo: "pattern", patrones: []patronXSD{patron002}},
}}
var tRetPlanesRetiro11_Planesderetiro_SistemaFinanc = &tipoXSD{nombre: "RetPlanesRetiro11 Planesderetiro/@SistemaFinanc", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPlanesRetiro11_Planesderetiro_SistemaFinanc},
}}
var tRetPlanesRetiro11_Planesderetiro_HuboRetirosAnioInmAntPer = &tipoXSD{nombre: "RetPlanesRetiro11 Planesderetiro/@HuboRetirosAnioInmAntPer", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPlanesRetiro11_Planesderetiro_HuboRetirosAnioInmAntPer},
}}
var tRetPlanesRetiro11_Planesderetiro_HuboRetirosAnioInmAnt = &tipoXSD{nombre: "RetPlanesRetiro11 Planesderetiro/@HuboRetirosAnioInmAnt", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catRetPlanesRetiro11_Planesderetiro_HuboRetirosAnioInmAnt},
}}
var tRetPlanesRetiro11_Planesderetiro_NumReferencia = &tipoXSD{nombre: "RetPlanesRetiro11 Planesderetiro/@NumReferencia", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tRetPremios_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tRetSectorFinanciero_SectorFinanciero_IdFideicom = &tipoXSD{nombre: "RetSectorFinanciero SectorFinanciero/@IdFideicom", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 20},
}}
var tRetSectorFinanciero_SectorFinanciero_NomFideicom = &tipoXSD{nombre: "RetSectorFinanciero SectorFinanciero/@NomFideicom", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 100},
}}
var tRetSectorFinanciero_SectorFinanciero_DescripFideicom = &tipoXSD{nombre: "RetSectorFinanciero SectorFinanciero/@DescripFideicom", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
	{tipo: "maxLength", n: 300},
}}
var tTFD10_TimbreFiscalDigital_UUID = &tipoXSD{nombre: "TFD10 TimbreFiscalDigital/@UUID", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 36},
	{tipo: "pattern", patrones: []patronXSD{patron001}},
}}
var tTFD10_TimbreFiscalDigital_FechaTimbrado = &tipoXSD{nombre: "TFD10 TimbreFiscalDigital/@FechaTimbrado", prim: "dateTime", ws: "collapse"}
var tTFD10_TimbreFiscalDigital_selloCFD = &tipoXSD{nombre: "TFD10 TimbreFiscalDigital/@selloCFD", prim: "string", ws: "collapse"}
var tTFD10_TimbreFiscalDigital_noCertificadoSAT = &tipoXSD{nombre: "TFD10 TimbreFiscalDigital/@noCertificadoSAT", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "length", n: 20},
}}
var tTFD10_TimbreFiscalDigital_selloSAT = &tipoXSD{nombre: "TFD10 TimbreFiscalDigital/@selloSAT", prim: "string", ws: "collapse"}
var tDonat10_Donatarias_leyenda = &tipoXSD{nombre: "Donat10 Donatarias/@leyenda", prim: "string", ws: "collapse"}
var tVentaVehiculos10_VentaVehiculos_InformacionAduanera_numero = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_VentaVehiculos_InformacionAduanera_fecha = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tVentaVehiculos10_VentaVehiculos_InformacionAduanera_aduana = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera_numero = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/InformacionAduanera/@numero", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera_fecha = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/InformacionAduanera/@fecha", prim: "date", ws: "collapse"}
var tVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera_aduana = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/InformacionAduanera/@aduana", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_VentaVehiculos_Parte_cantidad = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/@cantidad", prim: "decimal", ws: "collapse"}
var tVentaVehiculos10_VentaVehiculos_Parte_unidad = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/@unidad", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_VentaVehiculos_Parte_noIdentificacion = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/@noIdentificacion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_VentaVehiculos_Parte_descripcion = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/Parte/@descripcion", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tVentaVehiculos10_tImporte = &tipoXSD{nombre: "t_Importe", prim: "decimal", ws: "collapse", facetas: []facetaXSD{
	{tipo: "fractionDigits", n: 6},
}}
var tVentaVehiculos10_VentaVehiculos_ClaveVehicular = &tipoXSD{nombre: "VentaVehiculos10 VentaVehiculos/@ClaveVehicular", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "minLength", n: 1},
}}
var tINE10_INE_Entidad_Contabilidad_IdContabilidad = &tipoXSD{nombre: "INE10 INE/Entidad/Contabilidad/@IdContabilidad", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 6},
}}
var tINE10_tClaveEntidad = &tipoXSD{nombre: "t_ClaveEntidad", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tClaveEntidad},
}}
var tINE10_tAmbito = &tipoXSD{nombre: "t_Ambito", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tAmbito},
}}
var tINE10_INE_Entidad_Ambito = &tipoXSD{nombre: "INE10 INE/Entidad/@Ambito", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tAmbito},
}}
var tINE10_INE_Version = &tipoXSD{nombre: "INE10 INE/@Version", prim: "string", ws: "collapse"}
var tINE10_tTipoProc = &tipoXSD{nombre: "t_TipoProc", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tTipoProc},
}}
var tINE10_INE_TipoProceso = &tipoXSD{nombre: "INE10 INE/@TipoProceso", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tTipoProc},
}}
var tINE10_tTipoComite = &tipoXSD{nombre: "t_TipoComite", prim: "string", ws: "preserve", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tTipoComite},
}}
var tINE10_INE_TipoComite = &tipoXSD{nombre: "INE10 INE/@TipoComite", prim: "string", ws: "collapse", facetas: []facetaXSD{
	{tipo: "enumeration", catalogo: catINE10_tTipoComite},
}}
var tINE10_INE_IdContabilidad = &tipoXSD{nombre: "INE10 INE/@IdContabilidad", prim: "int", ws: "collapse", facetas: []facetaXSD{
	{tipo: "totalDigits", n: 6},
}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}InformacionGlobal
var eCFDI40_Comprobante_InformacionGlobal = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "InformacionGlobal", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Periodicidad", tipo: tCatCFDI_cPeriodicidad, requerido: true},
		{nombre: "Meses", tipo: tCatCFDI_cMeses, requerido: true},
		{nombre: "Año", tipo: tCFDI40_Comprobante_InformacionGlobal_Anio, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}CfdiRelacionado
var eCFDI40_Comprobante_CfdiRelacionados_CfdiRelacionado = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "CfdiRelacionado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "UUID", tipo: tCFDI40_Comprobante_CfdiRelacionados_CfdiRelacionado_UUID, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}CfdiRelacionados
var eCFDI40_Comprobante_CfdiRelacionados = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "CfdiRelacionados", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoRelacion", tipo: tCatCFDI_cTipoRelacion, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_CfdiRelacionados_CfdiRelacionado, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Emisor
var eCFDI40_Comprobante_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Emisor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Rfc", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "Nombre", tipo: tCFDI40_Comprobante_Emisor_Nombre, requerido: true},
		{nombre: "RegimenFiscal", tipo: tCatCFDI_cRegimenFiscal, requerido: true},
		{nombre: "FacAtrAdquirente", tipo: tCFDI40_Comprobante_Emisor_FacAtrAdquirente},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Receptor
var eCFDI40_Comprobante_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Receptor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Rfc", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "Nombre", tipo: tCFDI40_Comprobante_Receptor_Nombre, requerido: true},
		{nombre: "DomicilioFiscalReceptor", tipo: tCFDI40_Comprobante_Receptor_DomicilioFiscalReceptor, requerido: true},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumRegIdTrib", tipo: tCFDI40_Comprobante_Receptor_NumRegIdTrib},
		{nombre: "RegimenFiscalReceptor", tipo: tCatCFDI_cRegimenFiscal, requerido: true},
		{nombre: "UsoCFDI", tipo: tCatCFDI_cUsoCFDI, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Traslado
var eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_Base, requerido: true},
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_TasaOCuota},
		{nombre: "Importe", tipo: tTdCFDI_tImporte},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Traslados
var eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Retencion
var eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Retencion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_Base, requerido: true},
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_TasaOCuota, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Retenciones
var eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Impuestos
var eCFDI40_Comprobante_Conceptos_Concepto_Impuestos = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Impuestos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Traslados, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Impuestos_Retenciones, min: 0, max: 1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}ACuentaTerceros
var eCFDI40_Comprobante_Conceptos_Concepto_ACuentaTerceros = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "ACuentaTerceros", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RfcACuentaTerceros", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "NombreACuentaTerceros", tipo: tCFDI40_Comprobante_Conceptos_Concepto_ACuentaTerceros_NombreACuentaTerceros, requerido: true},
		{nombre: "RegimenFiscalACuentaTerceros", tipo: tCatCFDI_cRegimenFiscal, requerido: true},
		{nombre: "DomicilioFiscalACuentaTerceros", tipo: tCFDI40_Comprobante_Conceptos_Concepto_ACuentaTerceros_DomicilioFiscalACuentaTerceros, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}InformacionAduanera
var eCFDI40_Comprobante_Conceptos_Concepto_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroPedimento", tipo: tCFDI40_Comprobante_Conceptos_Concepto_InformacionAduanera_NumeroPedimento, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}CuentaPredial
var eCFDI40_Comprobante_Conceptos_Concepto_CuentaPredial = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "CuentaPredial", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Numero", tipo: tCFDI40_Comprobante_Conceptos_Concepto_CuentaPredial_Numero, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}ComplementoConcepto
var eCFDI40_Comprobante_Conceptos_Concepto_ComplementoConcepto = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "ComplementoConcepto", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}InformacionAduanera
var eCFDI40_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroPedimento", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera_NumeroPedimento, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Parte
var eCFDI40_Comprobante_Conceptos_Concepto_Parte = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Parte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ClaveProdServ", tipo: tCatCFDI_cClaveProdServ, requerido: true},
		{nombre: "NoIdentificacion", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Parte_NoIdentificacion},
		{nombre: "Cantidad", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Parte_Cantidad, requerido: true},
		{nombre: "Unidad", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Parte_Unidad},
		{nombre: "Descripcion", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Parte_Descripcion, requerido: true},
		{nombre: "ValorUnitario", tipo: tTdCFDI_tImporte},
		{nombre: "Importe", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera, min: 0, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Concepto
var eCFDI40_Comprobante_Conceptos_Concepto = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Concepto", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ClaveProdServ", tipo: tCatCFDI_cClaveProdServ, requerido: true},
		{nombre: "NoIdentificacion", tipo: tCFDI40_Comprobante_Conceptos_Concepto_NoIdentificacion},
		{nombre: "Cantidad", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Cantidad, requerido: true},
		{nombre: "ClaveUnidad", tipo: tCatCFDI_cClaveUnidad, requerido: true},
		{nombre: "Unidad", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Unidad},
		{nombre: "Descripcion", tipo: tCFDI40_Comprobante_Conceptos_Concepto_Descripcion, requerido: true},
		{nombre: "ValorUnitario", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Descuento", tipo: tTdCFDI_tImporte},
		{nombre: "ObjetoImp", tipo: tCatCFDI_cObjetoImp, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Impuestos, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_ACuentaTerceros, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_InformacionAduanera, min: 0, max: -1},
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_CuentaPredial, min: 0, max: -1},
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_ComplementoConcepto, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Conceptos_Concepto_Parte, min: 0, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Conceptos
var eCFDI40_Comprobante_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Conceptos_Concepto, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Retencion
var eCFDI40_Comprobante_Impuestos_Retenciones_Retencion = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Retencion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Retenciones
var eCFDI40_Comprobante_Impuestos_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Impuestos_Retenciones_Retencion, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Traslado
var eCFDI40_Comprobante_Impuestos_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tCFDI40_Comprobante_Impuestos_Traslados_Traslado_TasaOCuota},
		{nombre: "Importe", tipo: tTdCFDI_tImporte},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Traslados
var eCFDI40_Comprobante_Impuestos_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Impuestos_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Impuestos
var eCFDI40_Comprobante_Impuestos = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Impuestos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalImpuestosRetenidos", tipo: tTdCFDI_tImporte},
		{nombre: "TotalImpuestosTrasladados", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_Impuestos_Retenciones, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Impuestos_Traslados, min: 0, max: 1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Complemento
var eCFDI40_Comprobante_Complemento = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Complemento", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 0, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Addenda
var eCFDI40_Comprobante_Addenda = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Addenda", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 1, max: -1},
	}}

// cfd/4/cfdv40.xsd: {http://www.sat.gob.mx/cfd/4}Comprobante
var eCFDI40_Comprobante = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/4", nombre: "Comprobante", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tCFDI40_Comprobante_Version, requerido: true, fijo: "4.0", conFijo: true},
		{nombre: "Serie", tipo: tCFDI40_Comprobante_Serie},
		{nombre: "Folio", tipo: tCFDI40_Comprobante_Folio},
		{nombre: "Fecha", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "Sello", tipo: tCFDI40_Comprobante_Sello, requerido: true},
		{nombre: "FormaPago", tipo: tCatCFDI_cFormaPago},
		{nombre: "NoCertificado", tipo: tCFDI40_Comprobante_NoCertificado, requerido: true},
		{nombre: "Certificado", tipo: tCFDI40_Comprobante_Certificado, requerido: true},
		{nombre: "CondicionesDePago", tipo: tCFDI40_Comprobante_CondicionesDePago},
		{nombre: "SubTotal", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Descuento", tipo: tTdCFDI_tImporte},
		{nombre: "Moneda", tipo: tCatCFDI_cMoneda, requerido: true},
		{nombre: "TipoCambio", tipo: tCFDI40_Comprobante_TipoCambio},
		{nombre: "Total", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "TipoDeComprobante", tipo: tCatCFDI_cTipoDeComprobante, requerido: true},
		{nombre: "Exportacion", tipo: tCatCFDI_cExportacion, requerido: true},
		{nombre: "MetodoPago", tipo: tCatCFDI_cMetodoPago},
		{nombre: "LugarExpedicion", tipo: tCatCFDI_cCodigoPostal, requerido: true},
		{nombre: "Confirmacion", tipo: tCFDI40_Comprobante_Confirmacion},
	},
	particulas: []particulaXSD{
		{elem: eCFDI40_Comprobante_InformacionGlobal, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_CfdiRelacionados, min: 0, max: -1},
		{elem: eCFDI40_Comprobante_Emisor, min: 1, max: 1},
		{elem: eCFDI40_Comprobante_Receptor, min: 1, max: 1},
		{elem: eCFDI40_Comprobante_Conceptos, min: 1, max: 1},
		{elem: eCFDI40_Comprobante_Impuestos, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Complemento, min: 0, max: 1},
		{elem: eCFDI40_Comprobante_Addenda, min: 0, max: 1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}CfdiRelacionado
var eCFDI33_Comprobante_CfdiRelacionados_CfdiRelacionado = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "CfdiRelacionado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "UUID", tipo: tCFDI33_Comprobante_CfdiRelacionados_CfdiRelacionado_UUID, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}CfdiRelacionados
var eCFDI33_Comprobante_CfdiRelacionados = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "CfdiRelacionados", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoRelacion", tipo: tCatCFDI_cTipoRelacion, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_CfdiRelacionados_CfdiRelacionado, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Emisor
var eCFDI33_Comprobante_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Emisor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Rfc", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "Nombre", tipo: tCFDI33_Comprobante_Emisor_Nombre},
		{nombre: "RegimenFiscal", tipo: tCatCFDI_cRegimenFiscal, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Receptor
var eCFDI33_Comprobante_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Receptor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Rfc", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "Nombre", tipo: tCFDI33_Comprobante_Receptor_Nombre},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumRegIdTrib", tipo: tCFDI33_Comprobante_Receptor_NumRegIdTrib},
		{nombre: "UsoCFDI", tipo: tCatCFDI_cUsoCFDI, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Traslado
var eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_Base, requerido: true},
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado_TasaOCuota},
		{nombre: "Importe", tipo: tTdCFDI_tImporte},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Traslados
var eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Retencion
var eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Retencion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_Base, requerido: true},
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion_TasaOCuota, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Retenciones
var eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones_Retencion, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Impuestos
var eCFDI33_Comprobante_Conceptos_Concepto_Impuestos = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Impuestos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Traslados, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Impuestos_Retenciones, min: 0, max: 1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}InformacionAduanera
var eCFDI33_Comprobante_Conceptos_Concepto_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroPedimento", tipo: tCFDI33_Comprobante_Conceptos_Concepto_InformacionAduanera_NumeroPedimento, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}CuentaPredial
var eCFDI33_Comprobante_Conceptos_Concepto_CuentaPredial = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "CuentaPredial", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Numero", tipo: tCFDI33_Comprobante_Conceptos_Concepto_CuentaPredial_Numero, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}ComplementoConcepto
var eCFDI33_Comprobante_Conceptos_Concepto_ComplementoConcepto = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "ComplementoConcepto", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}InformacionAduanera
var eCFDI33_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroPedimento", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera_NumeroPedimento, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Parte
var eCFDI33_Comprobante_Conceptos_Concepto_Parte = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Parte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ClaveProdServ", tipo: tCatCFDI_cClaveProdServ, requerido: true},
		{nombre: "NoIdentificacion", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Parte_NoIdentificacion},
		{nombre: "Cantidad", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Parte_Cantidad, requerido: true},
		{nombre: "Unidad", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Parte_Unidad},
		{nombre: "Descripcion", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Parte_Descripcion, requerido: true},
		{nombre: "ValorUnitario", tipo: tTdCFDI_tImporte},
		{nombre: "Importe", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Parte_InformacionAduanera, min: 0, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Concepto
var eCFDI33_Comprobante_Conceptos_Concepto = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Concepto", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ClaveProdServ", tipo: tCatCFDI_cClaveProdServ, requerido: true},
		{nombre: "NoIdentificacion", tipo: tCFDI33_Comprobante_Conceptos_Concepto_NoIdentificacion},
		{nombre: "Cantidad", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Cantidad, requerido: true},
		{nombre: "ClaveUnidad", tipo: tCatCFDI_cClaveUnidad, requerido: true},
		{nombre: "Unidad", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Unidad},
		{nombre: "Descripcion", tipo: tCFDI33_Comprobante_Conceptos_Concepto_Descripcion, requerido: true},
		{nombre: "ValorUnitario", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Descuento", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Impuestos, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_InformacionAduanera, min: 0, max: -1},
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_CuentaPredial, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_ComplementoConcepto, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Conceptos_Concepto_Parte, min: 0, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Conceptos
var eCFDI33_Comprobante_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Conceptos_Concepto, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Retencion
var eCFDI33_Comprobante_Impuestos_Retenciones_Retencion = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Retencion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Retenciones
var eCFDI33_Comprobante_Impuestos_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Impuestos_Retenciones_Retencion, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Traslado
var eCFDI33_Comprobante_Impuestos_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tCFDI33_Comprobante_Impuestos_Traslados_Traslado_TasaOCuota, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Traslados
var eCFDI33_Comprobante_Impuestos_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Impuestos_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Impuestos
var eCFDI33_Comprobante_Impuestos = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Impuestos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalImpuestosRetenidos", tipo: tTdCFDI_tImporte},
		{nombre: "TotalImpuestosTrasladados", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_Impuestos_Retenciones, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Impuestos_Traslados, min: 0, max: 1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Complemento
var eCFDI33_Comprobante_Complemento = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Complemento", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 0, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Addenda
var eCFDI33_Comprobante_Addenda = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Addenda", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 1, max: -1},
	}}

// cfd/3/cfdv33.xsd: {http://www.sat.gob.mx/cfd/3}Comprobante
var eCFDI33_Comprobante = &elementoXSD{ns: "http://www.sat.gob.mx/cfd/3", nombre: "Comprobante", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tCFDI33_Comprobante_Version, requerido: true, fijo: "3.3", conFijo: true},
		{nombre: "Serie", tipo: tCFDI33_Comprobante_Serie},
		{nombre: "Folio", tipo: tCFDI33_Comprobante_Folio},
		{nombre: "Fecha", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "Sello", tipo: tCFDI33_Comprobante_Sello, requerido: true},
		{nombre: "FormaPago", tipo: tCatCFDI_cFormaPago},
		{nombre: "NoCertificado", tipo: tCFDI33_Comprobante_NoCertificado, requerido: true},
		{nombre: "Certificado", tipo: tCFDI33_Comprobante_Certificado, requerido: true},
		{nombre: "CondicionesDePago", tipo: tCFDI33_Comprobante_CondicionesDePago},
		{nombre: "SubTotal", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Descuento", tipo: tTdCFDI_tImporte},
		{nombre: "Moneda", tipo: tCatCFDI_cMoneda, requerido: true},
		{nombre: "TipoCambio", tipo: tCFDI33_Comprobante_TipoCambio},
		{nombre: "Total", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "TipoDeComprobante", tipo: tCatCFDI_cTipoDeComprobante, requerido: true},
		{nombre: "MetodoPago", tipo: tCatCFDI_cMetodoPago},
		{nombre: "LugarExpedicion", tipo: tCatCFDI_cCodigoPostal, requerido: true},
		{nombre: "Confirmacion", tipo: tCFDI33_Comprobante_Confirmacion},
	},
	particulas: []particulaXSD{
		{elem: eCFDI33_Comprobante_CfdiRelacionados, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Emisor, min: 1, max: 1},
		{elem: eCFDI33_Comprobante_Receptor, min: 1, max: 1},
		{elem: eCFDI33_Comprobante_Conceptos, min: 1, max: 1},
		{elem: eCFDI33_Comprobante_Impuestos, min: 0, max: 1},
		{elem: eCFDI33_Comprobante_Complemento, min: 0, max: -1},
		{elem: eCFDI33_Comprobante_Addenda, min: 0, max: 1},
	}}

// cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd: {http://www.sat.gob.mx/TimbreFiscalDigital}TimbreFiscalDigital
var eTFD11_TimbreFiscalDigital = &elementoXSD{ns: "http://www.sat.gob.mx/TimbreFiscalDigital", nombre: "TimbreFiscalDigital", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "UUID", tipo: tTFD11_TimbreFiscalDigital_UUID, requerido: true},
		{nombre: "FechaTimbrado", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "RfcProvCertif", tipo: tTdCFDI_tRFCPM, requerido: true},
		{nombre: "Leyenda", tipo: tTFD11_TimbreFiscalDigital_Leyenda},
		{nombre: "SelloCFD", tipo: tTFD11_TimbreFiscalDigital_SelloCFD, requerido: true},
		{nombre: "NoCertificadoSAT", tipo: tTFD11_TimbreFiscalDigital_NoCertificadoSAT, requerido: true},
		{nombre: "SelloSAT", tipo: tTFD11_TimbreFiscalDigital_SelloSAT, requerido: true},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}Totales
var ePagos20_Pagos_Totales = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "Totales", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TotalRetencionesIVA", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalRetencionesISR", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalRetencionesIEPS", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosBaseIVA16", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosImpuestoIVA16", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosBaseIVA8", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosImpuestoIVA8", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosBaseIVA0", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosImpuestoIVA0", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalTrasladosBaseIVAExento", tipo: tTdCFDI_tImporteMXN},
		{nombre: "MontoTotalPagos", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}RetencionDR
var ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR_RetencionDR = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "RetencionDR", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BaseDR", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR_RetencionDR_BaseDR, requerido: true},
		{nombre: "ImpuestoDR", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactorDR", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuotaDR", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR_RetencionDR_TasaOCuotaDR, requerido: true},
		{nombre: "ImporteDR", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}RetencionesDR
var ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "RetencionesDR", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR_RetencionDR, min: 1, max: -1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}TrasladoDR
var ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR_TrasladoDR = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "TrasladoDR", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BaseDR", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR_TrasladoDR_BaseDR, requerido: true},
		{nombre: "ImpuestoDR", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactorDR", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuotaDR", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR_TrasladoDR_TasaOCuotaDR},
		{nombre: "ImporteDR", tipo: tTdCFDI_tImporte},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}TrasladosDR
var ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "TrasladosDR", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR_TrasladoDR, min: 1, max: -1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}ImpuestosDR
var ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "ImpuestosDR", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_RetencionesDR, min: 0, max: 1},
		{elem: ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR_TrasladosDR, min: 0, max: 1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}DoctoRelacionado
var ePagos20_Pagos_Pago_DoctoRelacionado = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "DoctoRelacionado", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "IdDocumento", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_IdDocumento, requerido: true},
		{nombre: "Serie", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_Serie},
		{nombre: "Folio", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_Folio},
		{nombre: "MonedaDR", tipo: tCatCFDI_cMoneda, requerido: true},
		{nombre: "EquivalenciaDR", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_EquivalenciaDR},
		{nombre: "NumParcialidad", tipo: tPagos20_Pagos_Pago_DoctoRelacionado_NumParcialidad, requerido: true},
		{nombre: "ImpSaldoAnt", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "ImpPagado", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "ImpSaldoInsoluto", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "ObjetoImpDR", tipo: tCatCFDI_cObjetoImp, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_DoctoRelacionado_ImpuestosDR, min: 0, max: 1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}RetencionP
var ePagos20_Pagos_Pago_ImpuestosP_RetencionesP_RetencionP = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "RetencionP", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ImpuestoP", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "ImporteP", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}RetencionesP
var ePagos20_Pagos_Pago_ImpuestosP_RetencionesP = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "RetencionesP", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_ImpuestosP_RetencionesP_RetencionP, min: 1, max: -1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}TrasladoP
var ePagos20_Pagos_Pago_ImpuestosP_TrasladosP_TrasladoP = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "TrasladoP", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BaseP", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "ImpuestoP", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactorP", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuotaP", tipo: tPagos20_Pagos_Pago_ImpuestosP_TrasladosP_TrasladoP_TasaOCuotaP},
		{nombre: "ImporteP", tipo: tTdCFDI_tImporte},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}TrasladosP
var ePagos20_Pagos_Pago_ImpuestosP_TrasladosP = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "TrasladosP", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_ImpuestosP_TrasladosP_TrasladoP, min: 1, max: -1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}ImpuestosP
var ePagos20_Pagos_Pago_ImpuestosP = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "ImpuestosP", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_ImpuestosP_RetencionesP, min: 0, max: 1},
		{elem: ePagos20_Pagos_Pago_ImpuestosP_TrasladosP, min: 0, max: 1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}Pago
var ePagos20_Pagos_Pago = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "Pago", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "FechaPago", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "FormaDePagoP", tipo: tCatCFDI_cFormaPago, requerido: true},
		{nombre: "MonedaP", tipo: tCatCFDI_cMoneda, requerido: true},
		{nombre: "TipoCambioP", tipo: tPagos20_Pagos_Pago_TipoCambioP},
		{nombre: "Monto", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "NumOperacion", tipo: tPagos20_Pagos_Pago_NumOperacion},
		{nombre: "RfcEmisorCtaOrd", tipo: tPagos20_Pagos_Pago_RfcEmisorCtaOrd},
		{nombre: "NomBancoOrdExt", tipo: tPagos20_Pagos_Pago_NomBancoOrdExt},
		{nombre: "CtaOrdenante", tipo: tPagos20_Pagos_Pago_CtaOrdenante},
		{nombre: "RfcEmisorCtaBen", tipo: tTdCFDI_tRFCPM},
		{nombre: "CtaBeneficiario", tipo: tPagos20_Pagos_Pago_CtaBeneficiario},
		{nombre: "TipoCadPago", tipo: tCatPagos_cTipoCadenaPago},
		{nombre: "CertPago", tipo: tPagos20_Pagos_Pago_CertPago},
		{nombre: "CadPago", tipo: tPagos20_Pagos_Pago_CadPago},
		{nombre: "SelloPago", tipo: tPagos20_Pagos_Pago_SelloPago},
	},
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Pago_DoctoRelacionado, min: 1, max: -1},
		{elem: ePagos20_Pagos_Pago_ImpuestosP, min: 0, max: 1},
	}}

// cfd/Pagos/Pagos20.xsd: {http://www.sat.gob.mx/Pagos20}Pagos
var ePagos20_Pagos = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos20", nombre: "Pagos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tPagos20_Pagos_Version, requerido: true, fijo: "2.0", conFijo: true},
	},
	particulas: []particulaXSD{
		{elem: ePagos20_Pagos_Totales, min: 1, max: 1},
		{elem: ePagos20_Pagos_Pago, min: 1, max: -1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}EntidadSNCF
var eNomina12_Nomina_Emisor_EntidadSNCF = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "EntidadSNCF", vacio: true,
	attrs: []atributoXSD{
		{nombre: "OrigenRecurso", tipo: tCatNomina_cOrigenRecurso, requerido: true},
		{nombre: "MontoRecursoPropio", tipo: tTdCFDI_tImporteMXN},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Emisor
var eNomina12_Nomina_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Emisor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Curp", tipo: tTdCFDI_tCURP},
		{nombre: "RegistroPatronal", tipo: tNomina12_Nomina_Emisor_RegistroPatronal},
		{nombre: "RfcPatronOrigen", tipo: tTdCFDI_tRFC},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Emisor_EntidadSNCF, min: 0, max: 1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}SubContratacion
var eNomina12_Nomina_Receptor_SubContratacion = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "SubContratacion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RfcLabora", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "PorcentajeTiempo", tipo: tNomina12_Nomina_Receptor_SubContratacion_PorcentajeTiempo, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Receptor
var eNomina12_Nomina_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Receptor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Curp", tipo: tTdCFDI_tCURP, requerido: true},
		{nombre: "NumSeguridadSocial", tipo: tNomina12_Nomina_Receptor_NumSeguridadSocial},
		{nombre: "FechaInicioRelLaboral", tipo: tTdCFDI_tFecha},
		{nombre: "Antigüedad", tipo: tNomina12_Nomina_Receptor_Antiguedad},
		{nombre: "TipoContrato", tipo: tCatNomina_cTipoContrato, requerido: true},
		{nombre: "Sindicalizado", tipo: tNomina12_Nomina_Receptor_Sindicalizado},
		{nombre: "TipoJornada", tipo: tCatNomina_cTipoJornada},
		{nombre: "TipoRegimen", tipo: tCatNomina_cTipoRegimen, requerido: true},
		{nombre: "NumEmpleado", tipo: tNomina12_Nomina_Receptor_NumEmpleado, requerido: true},
		{nombre: "Departamento", tipo: tNomina12_Nomina_Receptor_Departamento},
		{nombre: "Puesto", tipo: tNomina12_Nomina_Receptor_Puesto},
		{nombre: "RiesgoPuesto", tipo: tCatNomina_cRiesgoPuesto},
		{nombre: "PeriodicidadPago", tipo: tCatNomina_cPeriodicidadPago, requerido: true},
		{nombre: "Banco", tipo: tCatNomina_cBanco},
		{nombre: "CuentaBancaria", tipo: tTdCFDI_tCuentaBancaria},
		{nombre: "SalarioBaseCotApor", tipo: tTdCFDI_tImporteMXN},
		{nombre: "SalarioDiarioIntegrado", tipo: tTdCFDI_tImporteMXN},
		{nombre: "ClaveEntFed", tipo: tCatCFDI_cEstado, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Receptor_SubContratacion, min: 0, max: -1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}AccionesOTitulos
var eNomina12_Nomina_Percepciones_Percepcion_AccionesOTitulos = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "AccionesOTitulos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ValorMercado", tipo: tNomina12_Nomina_Percepciones_Percepcion_AccionesOTitulos_ValorMercado, requerido: true},
		{nombre: "PrecioAlOtorgarse", tipo: tNomina12_Nomina_Percepciones_Percepcion_AccionesOTitulos_PrecioAlOtorgarse, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}HorasExtra
var eNomina12_Nomina_Percepciones_Percepcion_HorasExtra = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "HorasExtra", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Dias", tipo: tNomina12_Nomina_Percepciones_Percepcion_HorasExtra_Dias, requerido: true},
		{nombre: "TipoHoras", tipo: tCatNomina_cTipoHoras, requerido: true},
		{nombre: "HorasExtra", tipo: tNomina12_Nomina_Percepciones_Percepcion_HorasExtra_HorasExtra, requerido: true},
		{nombre: "ImportePagado", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Percepcion
var eNomina12_Nomina_Percepciones_Percepcion = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Percepcion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoPercepcion", tipo: tCatNomina_cTipoPercepcion, requerido: true},
		{nombre: "Clave", tipo: tNomina12_Nomina_Percepciones_Percepcion_Clave, requerido: true},
		{nombre: "Concepto", tipo: tNomina12_Nomina_Percepciones_Percepcion_Concepto, requerido: true},
		{nombre: "ImporteGravado", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "ImporteExento", tipo: tTdCFDI_tImporteMXN, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Percepciones_Percepcion_AccionesOTitulos, min: 0, max: 1},
		{elem: eNomina12_Nomina_Percepciones_Percepcion_HorasExtra, min: 0, max: -1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}JubilacionPensionRetiro
var eNomina12_Nomina_Percepciones_JubilacionPensionRetiro = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "JubilacionPensionRetiro", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TotalUnaExhibicion", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalParcialidad", tipo: tTdCFDI_tImporteMXN},
		{nombre: "MontoDiario", tipo: tTdCFDI_tImporteMXN},
		{nombre: "IngresoAcumulable", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "IngresoNoAcumulable", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}SeparacionIndemnizacion
var eNomina12_Nomina_Percepciones_SeparacionIndemnizacion = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "SeparacionIndemnizacion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TotalPagado", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "NumAñosServicio", tipo: tNomina12_Nomina_Percepciones_SeparacionIndemnizacion_NumAniosServicio, requerido: true},
		{nombre: "UltimoSueldoMensOrd", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "IngresoAcumulable", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "IngresoNoAcumulable", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Percepciones
var eNomina12_Nomina_Percepciones = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Percepciones", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalSueldos", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalSeparacionIndemnizacion", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalJubilacionPensionRetiro", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalGravado", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "TotalExento", tipo: tTdCFDI_tImporteMXN, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Percepciones_Percepcion, min: 1, max: -1},
		{elem: eNomina12_Nomina_Percepciones_JubilacionPensionRetiro, min: 0, max: 1},
		{elem: eNomina12_Nomina_Percepciones_SeparacionIndemnizacion, min: 0, max: 1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Deduccion
var eNomina12_Nomina_Deducciones_Deduccion = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Deduccion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDeduccion", tipo: tCatNomina_cTipoDeduccion, requerido: true},
		{nombre: "Clave", tipo: tNomina12_Nomina_Deducciones_Deduccion_Clave, requerido: true},
		{nombre: "Concepto", tipo: tNomina12_Nomina_Deducciones_Deduccion_Concepto, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Deducciones
var eNomina12_Nomina_Deducciones = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Deducciones", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalOtrasDeducciones", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalImpuestosRetenidos", tipo: tTdCFDI_tImporteMXN},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Deducciones_Deduccion, min: 1, max: -1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}SubsidioAlEmpleo
var eNomina12_Nomina_OtrosPagos_OtroPago_SubsidioAlEmpleo = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "SubsidioAlEmpleo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubsidioCausado", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}CompensacionSaldosAFavor
var eNomina12_Nomina_OtrosPagos_OtroPago_CompensacionSaldosAFavor = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "CompensacionSaldosAFavor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SaldoAFavor", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "Año", tipo: tNomina12_Nomina_OtrosPagos_OtroPago_CompensacionSaldosAFavor_Anio, requerido: true},
		{nombre: "RemanenteSalFav", tipo: tTdCFDI_tImporteMXN, requerido: true},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}OtroPago
var eNomina12_Nomina_OtrosPagos_OtroPago = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "OtroPago", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoOtroPago", tipo: tCatNomina_cTipoOtroPago, requerido: true},
		{nombre: "Clave", tipo: tNomina12_Nomina_OtrosPagos_OtroPago_Clave, requerido: true},
		{nombre: "Concepto", tipo: tNomina12_Nomina_OtrosPagos_OtroPago_Concepto, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporteMXN, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_OtrosPagos_OtroPago_SubsidioAlEmpleo, min: 0, max: 1},
		{elem: eNomina12_Nomina_OtrosPagos_OtroPago_CompensacionSaldosAFavor, min: 0, max: 1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}OtrosPagos
var eNomina12_Nomina_OtrosPagos = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "OtrosPagos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_OtrosPagos_OtroPago, min: 1, max: -1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Incapacidad
var eNomina12_Nomina_Incapacidades_Incapacidad = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Incapacidad", vacio: true,
	attrs: []atributoXSD{
		{nombre: "DiasIncapacidad", tipo: tNomina12_Nomina_Incapacidades_Incapacidad_DiasIncapacidad, requerido: true},
		{nombre: "TipoIncapacidad", tipo: tCatNomina_cTipoIncapacidad, requerido: true},
		{nombre: "ImporteMonetario", tipo: tTdCFDI_tImporteMXN},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Incapacidades
var eNomina12_Nomina_Incapacidades = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Incapacidades", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Incapacidades_Incapacidad, min: 1, max: -1},
	}}

// cfd/nomina/nomina12.xsd: {http://www.sat.gob.mx/nomina12}Nomina
var eNomina12_Nomina = &elementoXSD{ns: "http://www.sat.gob.mx/nomina12", nombre: "Nomina", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.2", conFijo: true},
		{nombre: "TipoNomina", tipo: tCatNomina_cTipoNomina, requerido: true},
		{nombre: "FechaPago", tipo: tTdCFDI_tFecha, requerido: true},
		{nombre: "FechaInicialPago", tipo: tTdCFDI_tFecha, requerido: true},
		{nombre: "FechaFinalPago", tipo: tTdCFDI_tFecha, requerido: true},
		{nombre: "NumDiasPagados", tipo: tNomina12_Nomina_NumDiasPagados, requerido: true},
		{nombre: "TotalPercepciones", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalDeducciones", tipo: tTdCFDI_tImporteMXN},
		{nombre: "TotalOtrosPagos", tipo: tTdCFDI_tImporteMXN},
	},
	particulas: []particulaXSD{
		{elem: eNomina12_Nomina_Emisor, min: 0, max: 1},
		{elem: eNomina12_Nomina_Receptor, min: 1, max: 1},
		{elem: eNomina12_Nomina_Percepciones, min: 0, max: 1},
		{elem: eNomina12_Nomina_Deducciones, min: 0, max: 1},
		{elem: eNomina12_Nomina_OtrosPagos, min: 0, max: 1},
		{elem: eNomina12_Nomina_Incapacidades, min: 0, max: 1},
	}}

// cfd/implocal/implocal.xsd: {http://www.sat.gob.mx/implocal}RetencionesLocales
var eImpLocal_ImpuestosLocales_RetencionesLocales = &elementoXSD{ns: "http://www.sat.gob.mx/implocal", nombre: "RetencionesLocales", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ImpLocRetenido", tipo: xsstring, requerido: true},
		{nombre: "TasadeRetencion", tipo: tImpLocal_ImpuestosLocales_RetencionesLocales_TasadeRetencion, requerido: true},
		{nombre: "Importe", tipo: tImpLocal_ImpuestosLocales_RetencionesLocales_Importe, requerido: true},
	}}

// cfd/implocal/implocal.xsd: {http://www.sat.gob.mx/implocal}TrasladosLocales
var eImpLocal_ImpuestosLocales_TrasladosLocales = &elementoXSD{ns: "http://www.sat.gob.mx/implocal", nombre: "TrasladosLocales", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ImpLocTrasladado", tipo: xsstring, requerido: true},
		{nombre: "TasadeTraslado", tipo: tImpLocal_ImpuestosLocales_TrasladosLocales_TasadeTraslado, requerido: true},
		{nombre: "Importe", tipo: tImpLocal_ImpuestosLocales_TrasladosLocales_Importe, requerido: true},
	}}

// cfd/implocal/implocal.xsd: {http://www.sat.gob.mx/implocal}ImpuestosLocales
var eImpLocal_ImpuestosLocales = &elementoXSD{ns: "http://www.sat.gob.mx/implocal", nombre: "ImpuestosLocales", seqMin: 1, seqMax: -1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TotaldeRetenciones", tipo: tImpLocal_ImpuestosLocales_TotaldeRetenciones, requerido: true},
		{nombre: "TotaldeTraslados", tipo: tImpLocal_ImpuestosLocales_TotaldeTraslados, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eImpLocal_ImpuestosLocales_RetencionesLocales, min: 0, max: 1},
		{elem: eImpLocal_ImpuestosLocales_TrasladosLocales, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}RegimenAduaneroCCP
var eCartaPorte31_CartaPorte_RegimenesAduaneros_RegimenAduaneroCCP = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "RegimenAduaneroCCP", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RegimenAduanero", tipo: tCatCartaPorte_cRegimenAduanero, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}RegimenesAduaneros
var eCartaPorte31_CartaPorte_RegimenesAduaneros = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "RegimenesAduaneros", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_RegimenesAduaneros_RegimenAduaneroCCP, min: 1, max: 10},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Domicilio
var eCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Ubicacion
var eCartaPorte31_CartaPorte_Ubicaciones_Ubicacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Ubicacion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoUbicacion", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion, requerido: true},
		{nombre: "IDUbicacion", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_IDUbicacion},
		{nombre: "RFCRemitenteDestinatario", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "NombreRemitenteDestinatario", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NombreRemitenteDestinatario},
		{nombre: "NumRegIdTrib", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NumRegIdTrib},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumEstacion", tipo: tCatCartaPorte_cEstaciones},
		{nombre: "NombreEstacion", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NombreEstacion},
		{nombre: "NavegacionTrafico", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico},
		{nombre: "FechaHoraSalidaLlegada", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "TipoEstacion", tipo: tCatCartaPorte_cTipoEstacion},
		{nombre: "DistanciaRecorrida", tipo: tCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Ubicaciones_Ubicacion_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Ubicaciones
var eCartaPorte31_CartaPorte_Ubicaciones = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Ubicaciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Ubicaciones_Ubicacion, min: 2, max: -1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}DocumentacionAduanera
var eCartaPorte31_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "DocumentacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDocumento", tipo: tCatCartaPorte_cDocumentoAduanero, requerido: true},
		{nombre: "NumPedimento", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_NumPedimento},
		{nombre: "IdentDocAduanero", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_IdentDocAduanero},
		{nombre: "RFCImpo", tipo: tTdCFDI_tRFC},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}GuiasIdentificacion
var eCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "GuiasIdentificacion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroGuiaIdentificacion", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_NumeroGuiaIdentificacion, requerido: true},
		{nombre: "DescripGuiaIdentificacion", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_DescripGuiaIdentificacion, requerido: true},
		{nombre: "PesoGuiaIdentificacion", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_PesoGuiaIdentificacion, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}CantidadTransporta
var eCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "CantidadTransporta", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Cantidad", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad, requerido: true},
		{nombre: "IDOrigen", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen, requerido: true},
		{nombre: "IDDestino", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino, requerido: true},
		{nombre: "CvesTransporte", tipo: tCatCartaPorte_cCveTransporte},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}DetalleMercancia
var eCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "DetalleMercancia", vacio: true,
	attrs: []atributoXSD{
		{nombre: "UnidadPesoMerc", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoBruto", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto, requerido: true},
		{nombre: "PesoNeto", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto, requerido: true},
		{nombre: "PesoTara", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara, requerido: true},
		{nombre: "NumPiezas", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Mercancia
var eCartaPorte31_CartaPorte_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "BienesTransp", tipo: tCatCartaPorte_cClaveProdServCP, requerido: true},
		{nombre: "ClaveSTCC", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_ClaveSTCC},
		{nombre: "Descripcion", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_Descripcion, requerido: true},
		{nombre: "Cantidad", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_Cantidad, requerido: true},
		{nombre: "ClaveUnidad", tipo: tCatCFDI_cClaveUnidad, requerido: true},
		{nombre: "Unidad", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_Unidad},
		{nombre: "Dimensiones", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_Dimensiones},
		{nombre: "MaterialPeligroso", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
		{nombre: "CveMaterialPeligroso", tipo: tCatCartaPorte_cMaterialPeligroso},
		{nombre: "Embalaje", tipo: tCatCartaPorte_cTipoEmbalaje},
		{nombre: "DescripEmbalaje", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DescripEmbalaje},
		{nombre: "SectorCOFEPRIS", tipo: tCatCartaPorte_cSectorCOFEPRIS},
		{nombre: "NombreIngredienteActivo", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_NombreIngredienteActivo},
		{nombre: "NomQuimico", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_NomQuimico},
		{nombre: "DenominacionGenericaProd", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DenominacionGenericaProd},
		{nombre: "DenominacionDistintivaProd", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DenominacionDistintivaProd},
		{nombre: "Fabricante", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_Fabricante},
		{nombre: "FechaCaducidad", tipo: tTdCFDI_tFecha},
		{nombre: "LoteMedicamento", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_LoteMedicamento},
		{nombre: "FormaFarmaceutica", tipo: tCatCartaPorte_cFormaFarmaceutica},
		{nombre: "CondicionesEspTransp", tipo: tCatCartaPorte_cCondicionesEspeciales},
		{nombre: "RegistroSanitarioFolioAutorizacion", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_RegistroSanitarioFolioAutorizacion},
		{nombre: "PermisoImportacion", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_PermisoImportacion},
		{nombre: "FolioImpoVUCEM", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_FolioImpoVUCEM},
		{nombre: "NumCAS", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_NumCAS},
		{nombre: "RazonSocialEmpImp", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_RazonSocialEmpImp},
		{nombre: "NumRegSanPlagCOFEPRIS", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_NumRegSanPlagCOFEPRIS},
		{nombre: "DatosFabricante", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DatosFabricante},
		{nombre: "DatosFormulador", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DatosFormulador},
		{nombre: "DatosMaquilador", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DatosMaquilador},
		{nombre: "UsoAutorizado", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_UsoAutorizado},
		{nombre: "PesoEnKg", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_PesoEnKg, requerido: true},
		{nombre: "ValorMercancia", tipo: tTdCFDI_tImporte},
		{nombre: "Moneda", tipo: tCatCFDI_cMoneda},
		{nombre: "FraccionArancelaria", tipo: tCatComExt_cFraccionArancelaria},
		{nombre: "UUIDComercioExt", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_UUIDComercioExt},
		{nombre: "TipoMateria", tipo: tCatCartaPorte_cTipoMateria},
		{nombre: "DescripcionMateria", tipo: tCartaPorte31_CartaPorte_Mercancias_Mercancia_DescripcionMateria},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera, min: 0, max: -1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion, min: 0, max: -1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_Mercancia_CantidadTransporta, min: 0, max: -1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_Mercancia_DetalleMercancia, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}IdentificacionVehicular
var eCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "IdentificacionVehicular", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ConfigVehicular", tipo: tCatCartaPorte_cConfigAutotransporte, requerido: true},
		{nombre: "PesoBrutoVehicular", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PesoBrutoVehicular, requerido: true},
		{nombre: "PlacaVM", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PlacaVM, requerido: true},
		{nombre: "AnioModeloVM", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_AnioModeloVM, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Seguros
var eCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Seguros", vacio: true,
	attrs: []atributoXSD{
		{nombre: "AseguraRespCivil", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraRespCivil, requerido: true},
		{nombre: "PolizaRespCivil", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaRespCivil, requerido: true},
		{nombre: "AseguraMedAmbiente", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraMedAmbiente},
		{nombre: "PolizaMedAmbiente", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaMedAmbiente},
		{nombre: "AseguraCarga", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraCarga},
		{nombre: "PolizaCarga", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaCarga},
		{nombre: "PrimaSeguro", tipo: tTdCFDI_tImporte},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Remolque
var eCartaPorte31_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Remolque", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubTipoRem", tipo: tCatCartaPorte_cSubTipoRem, requerido: true},
		{nombre: "Placa", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque_Placa, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Remolques
var eCartaPorte31_CartaPorte_Mercancias_Autotransporte_Remolques = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Remolques", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque, min: 1, max: 2},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Autotransporte
var eCartaPorte31_CartaPorte_Mercancias_Autotransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Autotransporte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte31_CartaPorte_Mercancias_Autotransporte_NumPermisoSCT, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular, min: 1, max: 1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_Autotransporte_Seguros, min: 1, max: 1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_Autotransporte_Remolques, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}RemolqueCCP
var eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_RemolquesCCP_RemolqueCCP = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "RemolqueCCP", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubTipoRemCCP", tipo: tCatCartaPorte_cSubTipoRem, requerido: true},
		{nombre: "PlacaCCP", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_RemolquesCCP_RemolqueCCP_PlacaCCP, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}RemolquesCCP
var eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_RemolquesCCP = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "RemolquesCCP", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_RemolquesCCP_RemolqueCCP, min: 1, max: 2},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Contenedor
var eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Contenedor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedorMaritimo, requerido: true},
		{nombre: "MatriculaContenedor", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor},
		{nombre: "NumPrecinto", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto},
		{nombre: "IdCCPRelacionado", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_IdCCPRelacionado},
		{nombre: "PlacaVMCCP", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_PlacaVMCCP},
		{nombre: "FechaCertificacionCCP", tipo: tTdCFDI_tFechaH},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_RemolquesCCP, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}TransporteMaritimo
var eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "TransporteMaritimo", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT},
		{nombre: "NombreAseg", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro},
		{nombre: "TipoEmbarcacion", tipo: tCatCartaPorte_cConfigMaritima, requerido: true},
		{nombre: "Matricula", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Matricula, requerido: true},
		{nombre: "NumeroOMI", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI, requerido: true},
		{nombre: "AnioEmbarcacion", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion},
		{nombre: "NombreEmbarc", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc},
		{nombre: "NacionalidadEmbarc", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "UnidadesDeArqBruto", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto, requerido: true},
		{nombre: "TipoCarga", tipo: tCatCartaPorte_cClaveTipoCarga, requerido: true},
		{nombre: "Eslora", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Eslora},
		{nombre: "Manga", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Manga},
		{nombre: "Calado", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Calado},
		{nombre: "Puntal", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Puntal},
		{nombre: "LineaNaviera", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera},
		{nombre: "NombreAgenteNaviero", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero, requerido: true},
		{nombre: "NumAutorizacionNaviero", tipo: tCatCartaPorte_cNumAutorizacionNaviero, requerido: true},
		{nombre: "NumViaje", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumViaje},
		{nombre: "NumConocEmbarc", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc},
		{nombre: "PermisoTempNavegacion", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_PermisoTempNavegacion},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo_Contenedor, min: 0, max: -1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}TransporteAereo
var eCartaPorte31_CartaPorte_Mercancias_TransporteAereo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "TransporteAereo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT, requerido: true},
		{nombre: "MatriculaAeronave", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave},
		{nombre: "NombreAseg", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro},
		{nombre: "NumeroGuia", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumeroGuia, requerido: true},
		{nombre: "LugarContrato", tipo: tTdCFDI_tDescrip120},
		{nombre: "CodigoTransportista", tipo: tCatCartaPorte_cCodigoTransporteAereo, requerido: true},
		{nombre: "RFCEmbarcador", tipo: tTdCFDI_tRFC},
		{nombre: "NumRegIdTribEmbarc", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc},
		{nombre: "ResidenciaFiscalEmbarc", tipo: tCatCFDI_cPais},
		{nombre: "NombreEmbarcador", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}DerechosDePaso
var eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "DerechosDePaso", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDerechoDePaso", tipo: tCatCartaPorte_cDerechosDePaso, requerido: true},
		{nombre: "KilometrajePagado", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Contenedor
var eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedor, requerido: true},
		{nombre: "PesoContenedorVacio", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio, requerido: true},
		{nombre: "PesoNetoMercancia", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Carro
var eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Carro", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoCarro", tipo: tCatCartaPorte_cTipoCarro, requerido: true},
		{nombre: "MatriculaCarro", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro, requerido: true},
		{nombre: "GuiaCarro", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro, requerido: true},
		{nombre: "ToneladasNetasCarro", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor, min: 0, max: -1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}TransporteFerroviario
var eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "TransporteFerroviario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoDeServicio", tipo: tCatCartaPorte_cTipoDeServicio, requerido: true},
		{nombre: "TipoDeTrafico", tipo: tCatCartaPorte_cTipoDeTrafico, requerido: true},
		{nombre: "NombreAseg", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso, min: 0, max: -1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario_Carro, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Mercancias
var eCartaPorte31_CartaPorte_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PesoBrutoTotal", tipo: tCartaPorte31_CartaPorte_Mercancias_PesoBrutoTotal, requerido: true},
		{nombre: "UnidadPeso", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoNetoTotal", tipo: tCartaPorte31_CartaPorte_Mercancias_PesoNetoTotal},
		{nombre: "NumTotalMercancias", tipo: tCartaPorte31_CartaPorte_Mercancias_NumTotalMercancias, requerido: true},
		{nombre: "CargoPorTasacion", tipo: tTdCFDI_tImporte},
		{nombre: "LogisticaInversaRecoleccionDevolucion", tipo: tCartaPorte31_CartaPorte_Mercancias_LogisticaInversaRecoleccionDevolucion},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_Mercancias_Mercancia, min: 1, max: -1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_Autotransporte, min: 0, max: 1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteMaritimo, min: 0, max: 1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteAereo, min: 0, max: 1},
		{elem: eCartaPorte31_CartaPorte_Mercancias_TransporteFerroviario, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}PartesTransporte
var eCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_PartesTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "PartesTransporte", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ParteTransporte", tipo: tCatCartaPorte_cParteTransporte, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}Domicilio
var eCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}TiposFigura
var eCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "TiposFigura", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoFigura", tipo: tCatCartaPorte_cFiguraTransporte, requerido: true},
		{nombre: "RFCFigura", tipo: tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_RFCFigura},
		{nombre: "NumLicencia", tipo: tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_NumLicencia},
		{nombre: "NombreFigura", tipo: tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_NombreFigura, requerido: true},
		{nombre: "NumRegIdTribFigura", tipo: tCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_NumRegIdTribFigura},
		{nombre: "ResidenciaFiscalFigura", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_PartesTransporte, min: 0, max: -1},
		{elem: eCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}FiguraTransporte
var eCartaPorte31_CartaPorte_FiguraTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "FiguraTransporte", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_FiguraTransporte_TiposFigura, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte31.xsd: {http://www.sat.gob.mx/CartaPorte31}CartaPorte
var eCartaPorte31_CartaPorte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte31", nombre: "CartaPorte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "3.1", conFijo: true},
		{nombre: "IdCCP", tipo: tCartaPorte31_CartaPorte_IdCCP, requerido: true},
		{nombre: "TranspInternac", tipo: tCartaPorte31_CartaPorte_TranspInternac, requerido: true},
		{nombre: "EntradaSalidaMerc", tipo: tCartaPorte31_CartaPorte_EntradaSalidaMerc},
		{nombre: "PaisOrigenDestino", tipo: tCatCFDI_cPais},
		{nombre: "ViaEntradaSalida", tipo: tCatCartaPorte_cCveTransporte},
		{nombre: "TotalDistRec", tipo: tCartaPorte31_CartaPorte_TotalDistRec},
		{nombre: "RegistroISTMO", tipo: tCartaPorte31_CartaPorte_RegistroISTMO},
		{nombre: "UbicacionPoloOrigen", tipo: tCatCartaPorte_cRegistroISTMO},
		{nombre: "UbicacionPoloDestino", tipo: tCatCartaPorte_cRegistroISTMO},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte31_CartaPorte_RegimenesAduaneros, min: 0, max: 1},
		{elem: eCartaPorte31_CartaPorte_Ubicaciones, min: 1, max: 1},
		{elem: eCartaPorte31_CartaPorte_Mercancias, min: 1, max: 1},
		{elem: eCartaPorte31_CartaPorte_FiguraTransporte, min: 0, max: 1},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}DoctoRelacionado
var ePagos10_Pagos_Pago_DoctoRelacionado = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "DoctoRelacionado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "IdDocumento", tipo: tPagos10_Pagos_Pago_DoctoRelacionado_IdDocumento, requerido: true},
		{nombre: "Serie", tipo: tPagos10_Pagos_Pago_DoctoRelacionado_Serie},
		{nombre: "Folio", tipo: tPagos10_Pagos_Pago_DoctoRelacionado_Folio},
		{nombre: "MonedaDR", tipo: tCatCFDI_cMoneda, requerido: true},
		{nombre: "TipoCambioDR", tipo: tPagos10_Pagos_Pago_DoctoRelacionado_TipoCambioDR},
		{nombre: "MetodoDePagoDR", tipo: tCatCFDI_cMetodoPago, requerido: true},
		{nombre: "NumParcialidad", tipo: tPagos10_Pagos_Pago_DoctoRelacionado_NumParcialidad},
		{nombre: "ImpSaldoAnt", tipo: tTdCFDI_tImporte},
		{nombre: "ImpPagado", tipo: tTdCFDI_tImporte},
		{nombre: "ImpSaldoInsoluto", tipo: tTdCFDI_tImporte},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Retencion
var ePagos10_Pagos_Pago_Impuestos_Retenciones_Retencion = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Retencion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Retenciones
var ePagos10_Pagos_Pago_Impuestos_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos10_Pagos_Pago_Impuestos_Retenciones_Retencion, min: 1, max: -1},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Traslado
var ePagos10_Pagos_Pago_Impuestos_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tCatCFDI_cImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: tCatCFDI_cTipoFactor, requerido: true},
		{nombre: "TasaOCuota", tipo: tPagos10_Pagos_Pago_Impuestos_Traslados_Traslado_TasaOCuota, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Traslados
var ePagos10_Pagos_Pago_Impuestos_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: ePagos10_Pagos_Pago_Impuestos_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Impuestos
var ePagos10_Pagos_Pago_Impuestos = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Impuestos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalImpuestosRetenidos", tipo: tTdCFDI_tImporte},
		{nombre: "TotalImpuestosTrasladados", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: ePagos10_Pagos_Pago_Impuestos_Retenciones, min: 0, max: 1},
		{elem: ePagos10_Pagos_Pago_Impuestos_Traslados, min: 0, max: 1},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Pago
var ePagos10_Pagos_Pago = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Pago", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "FechaPago", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "FormaDePagoP", tipo: tCatCFDI_cFormaPago, requerido: true},
		{nombre: "MonedaP", tipo: tCatCFDI_cMoneda, requerido: true},
		{nombre: "TipoCambioP", tipo: tPagos10_Pagos_Pago_TipoCambioP},
		{nombre: "Monto", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "NumOperacion", tipo: tPagos10_Pagos_Pago_NumOperacion},
		{nombre: "RfcEmisorCtaOrd", tipo: tPagos10_Pagos_Pago_RfcEmisorCtaOrd},
		{nombre: "NomBancoOrdExt", tipo: tPagos10_Pagos_Pago_NomBancoOrdExt},
		{nombre: "CtaOrdenante", tipo: tPagos10_Pagos_Pago_CtaOrdenante},
		{nombre: "RfcEmisorCtaBen", tipo: tTdCFDI_tRFCPM},
		{nombre: "CtaBeneficiario", tipo: tPagos10_Pagos_Pago_CtaBeneficiario},
		{nombre: "TipoCadPago", tipo: tCatPagos_cTipoCadenaPago},
		{nombre: "CertPago", tipo: tPagos10_Pagos_Pago_CertPago},
		{nombre: "CadPago", tipo: tPagos10_Pagos_Pago_CadPago},
		{nombre: "SelloPago", tipo: tPagos10_Pagos_Pago_SelloPago},
	},
	particulas: []particulaXSD{
		{elem: ePagos10_Pagos_Pago_DoctoRelacionado, min: 0, max: -1},
		{elem: ePagos10_Pagos_Pago_Impuestos, min: 0, max: -1},
	}}

// cfd/Pagos/Pagos10.xsd: {http://www.sat.gob.mx/Pagos}Pagos
var ePagos10_Pagos = &elementoXSD{ns: "http://www.sat.gob.mx/Pagos", nombre: "Pagos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tPagos10_Pagos_Version, requerido: true, fijo: "1.0", conFijo: true},
	},
	particulas: []particulaXSD{
		{elem: ePagos10_Pagos_Pago, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Domicilio
var eCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Ubicacion
var eCartaPorte30_CartaPorte_Ubicaciones_Ubicacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Ubicacion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoUbicacion", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion, requerido: true},
		{nombre: "IDUbicacion", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_IDUbicacion},
		{nombre: "RFCRemitenteDestinatario", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "NombreRemitenteDestinatario", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NombreRemitenteDestinatario},
		{nombre: "NumRegIdTrib", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NumRegIdTrib},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumEstacion", tipo: tCatCartaPorte_cEstaciones},
		{nombre: "NombreEstacion", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NombreEstacion},
		{nombre: "NavegacionTrafico", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico},
		{nombre: "FechaHoraSalidaLlegada", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "TipoEstacion", tipo: tCatCartaPorte_cTipoEstacion},
		{nombre: "DistanciaRecorrida", tipo: tCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Ubicaciones_Ubicacion_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Ubicaciones
var eCartaPorte30_CartaPorte_Ubicaciones = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Ubicaciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Ubicaciones_Ubicacion, min: 2, max: -1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}DocumentacionAduanera
var eCartaPorte30_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "DocumentacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDocumento", tipo: tCatCartaPorte_cDocumentoAduanero, requerido: true},
		{nombre: "NumPedimento", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_NumPedimento},
		{nombre: "IdentDocAduanero", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera_IdentDocAduanero},
		{nombre: "RFCImpo", tipo: tTdCFDI_tRFC},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}GuiasIdentificacion
var eCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "GuiasIdentificacion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroGuiaIdentificacion", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_NumeroGuiaIdentificacion, requerido: true},
		{nombre: "DescripGuiaIdentificacion", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_DescripGuiaIdentificacion, requerido: true},
		{nombre: "PesoGuiaIdentificacion", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_PesoGuiaIdentificacion, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}CantidadTransporta
var eCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "CantidadTransporta", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Cantidad", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad, requerido: true},
		{nombre: "IDOrigen", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen, requerido: true},
		{nombre: "IDDestino", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino, requerido: true},
		{nombre: "CvesTransporte", tipo: tCatCartaPorte_cCveTransporte},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}DetalleMercancia
var eCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "DetalleMercancia", vacio: true,
	attrs: []atributoXSD{
		{nombre: "UnidadPesoMerc", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoBruto", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto, requerido: true},
		{nombre: "PesoNeto", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto, requerido: true},
		{nombre: "PesoTara", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara, requerido: true},
		{nombre: "NumPiezas", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Mercancia
var eCartaPorte30_CartaPorte_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "BienesTransp", tipo: tCatCartaPorte_cClaveProdServCP, requerido: true},
		{nombre: "ClaveSTCC", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_ClaveSTCC},
		{nombre: "Descripcion", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_Descripcion, requerido: true},
		{nombre: "Cantidad", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_Cantidad, requerido: true},
		{nombre: "ClaveUnidad", tipo: tCatCFDI_cClaveUnidad, requerido: true},
		{nombre: "Unidad", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_Unidad},
		{nombre: "Dimensiones", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_Dimensiones},
		{nombre: "MaterialPeligroso", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
		{nombre: "CveMaterialPeligroso", tipo: tCatCartaPorte_cMaterialPeligroso},
		{nombre: "Embalaje", tipo: tCatCartaPorte_cTipoEmbalaje},
		{nombre: "DescripEmbalaje", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DescripEmbalaje},
		{nombre: "SectorCOFEPRIS", tipo: tCatCartaPorte_cSectorCOFEPRIS},
		{nombre: "NombreIngredienteActivo", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_NombreIngredienteActivo},
		{nombre: "NomQuimico", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_NomQuimico},
		{nombre: "DenominacionGenericaProd", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DenominacionGenericaProd},
		{nombre: "DenominacionDistintivaProd", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DenominacionDistintivaProd},
		{nombre: "Fabricante", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_Fabricante},
		{nombre: "FechaCaducidad", tipo: tTdCFDI_tFecha},
		{nombre: "LoteMedicamento", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_LoteMedicamento},
		{nombre: "FormaFarmaceutica", tipo: tCatCartaPorte_cFormaFarmaceutica},
		{nombre: "CondicionesEspTransp", tipo: tCatCartaPorte_cCondicionesEspeciales},
		{nombre: "RegistroSanitarioFolioAutorizacion", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_RegistroSanitarioFolioAutorizacion},
		{nombre: "PermisoImportacion", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_PermisoImportacion},
		{nombre: "FolioImpoVUCEM", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_FolioImpoVUCEM},
		{nombre: "NumCAS", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_NumCAS},
		{nombre: "RazonSocialEmpImp", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_RazonSocialEmpImp},
		{nombre: "NumRegSanPlagCOFEPRIS", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_NumRegSanPlagCOFEPRIS},
		{nombre: "DatosFabricante", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DatosFabricante},
		{nombre: "DatosFormulador", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DatosFormulador},
		{nombre: "DatosMaquilador", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DatosMaquilador},
		{nombre: "UsoAutorizado", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_UsoAutorizado},
		{nombre: "PesoEnKg", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_PesoEnKg, requerido: true},
		{nombre: "ValorMercancia", tipo: tTdCFDI_tImporte},
		{nombre: "Moneda", tipo: tCatCFDI_cMoneda},
		{nombre: "FraccionArancelaria", tipo: tCatComExt_cFraccionArancelaria},
		{nombre: "UUIDComercioExt", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_UUIDComercioExt},
		{nombre: "TipoMateria", tipo: tCatCartaPorte_cTipoMateria},
		{nombre: "DescripcionMateria", tipo: tCartaPorte30_CartaPorte_Mercancias_Mercancia_DescripcionMateria},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_Mercancia_DocumentacionAduanera, min: 0, max: -1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion, min: 0, max: -1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_Mercancia_CantidadTransporta, min: 0, max: -1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_Mercancia_DetalleMercancia, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}IdentificacionVehicular
var eCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "IdentificacionVehicular", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ConfigVehicular", tipo: tCatCartaPorte_cConfigAutotransporte, requerido: true},
		{nombre: "PesoBrutoVehicular", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PesoBrutoVehicular, requerido: true},
		{nombre: "PlacaVM", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PlacaVM, requerido: true},
		{nombre: "AnioModeloVM", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_AnioModeloVM, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Seguros
var eCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Seguros", vacio: true,
	attrs: []atributoXSD{
		{nombre: "AseguraRespCivil", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraRespCivil, requerido: true},
		{nombre: "PolizaRespCivil", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaRespCivil, requerido: true},
		{nombre: "AseguraMedAmbiente", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraMedAmbiente},
		{nombre: "PolizaMedAmbiente", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaMedAmbiente},
		{nombre: "AseguraCarga", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraCarga},
		{nombre: "PolizaCarga", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaCarga},
		{nombre: "PrimaSeguro", tipo: tTdCFDI_tImporte},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Remolque
var eCartaPorte30_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Remolque", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubTipoRem", tipo: tCatCartaPorte_cSubTipoRem, requerido: true},
		{nombre: "Placa", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque_Placa, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Remolques
var eCartaPorte30_CartaPorte_Mercancias_Autotransporte_Remolques = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Remolques", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque, min: 1, max: 2},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Autotransporte
var eCartaPorte30_CartaPorte_Mercancias_Autotransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Autotransporte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte30_CartaPorte_Mercancias_Autotransporte_NumPermisoSCT, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular, min: 1, max: 1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_Autotransporte_Seguros, min: 1, max: 1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_Autotransporte_Remolques, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Contenedor
var eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedorMaritimo, requerido: true},
		{nombre: "MatriculaContenedor", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor},
		{nombre: "NumPrecinto", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto},
		{nombre: "IdCCPRelacionado", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_IdCCPRelacionado},
		{nombre: "PlacaVMCCP", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_PlacaVMCCP},
		{nombre: "FechaCertificacionCCP", tipo: tTdCFDI_tFechaH},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}RemolqueCCP
var eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_RemolquesCCP_RemolqueCCP = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "RemolqueCCP", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubTipoRemCCP", tipo: tCatCartaPorte_cSubTipoRem, requerido: true},
		{nombre: "PlacaCCP", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_RemolquesCCP_RemolqueCCP_PlacaCCP, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}RemolquesCCP
var eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_RemolquesCCP = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "RemolquesCCP", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_RemolquesCCP_RemolqueCCP, min: 1, max: 2},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}TransporteMaritimo
var eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "TransporteMaritimo", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT},
		{nombre: "NombreAseg", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro},
		{nombre: "TipoEmbarcacion", tipo: tCatCartaPorte_cConfigMaritima, requerido: true},
		{nombre: "Matricula", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Matricula, requerido: true},
		{nombre: "NumeroOMI", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI, requerido: true},
		{nombre: "AnioEmbarcacion", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion},
		{nombre: "NombreEmbarc", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc},
		{nombre: "NacionalidadEmbarc", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "UnidadesDeArqBruto", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto, requerido: true},
		{nombre: "TipoCarga", tipo: tCatCartaPorte_cClaveTipoCarga, requerido: true},
		{nombre: "Eslora", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Eslora},
		{nombre: "Manga", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Manga},
		{nombre: "Calado", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Calado},
		{nombre: "Puntal", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Puntal},
		{nombre: "LineaNaviera", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera},
		{nombre: "NombreAgenteNaviero", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero, requerido: true},
		{nombre: "NumAutorizacionNaviero", tipo: tCatCartaPorte_cNumAutorizacionNaviero, requerido: true},
		{nombre: "NumViaje", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumViaje},
		{nombre: "NumConocEmbarc", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc},
		{nombre: "PermisoTempNavegacion", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_PermisoTempNavegacion},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_Contenedor, min: 0, max: -1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo_RemolquesCCP, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}TransporteAereo
var eCartaPorte30_CartaPorte_Mercancias_TransporteAereo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "TransporteAereo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT, requerido: true},
		{nombre: "MatriculaAeronave", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave},
		{nombre: "NombreAseg", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro},
		{nombre: "NumeroGuia", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumeroGuia, requerido: true},
		{nombre: "LugarContrato", tipo: tTdCFDI_tDescrip120},
		{nombre: "CodigoTransportista", tipo: tCatCartaPorte_cCodigoTransporteAereo, requerido: true},
		{nombre: "RFCEmbarcador", tipo: tTdCFDI_tRFC},
		{nombre: "NumRegIdTribEmbarc", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc},
		{nombre: "ResidenciaFiscalEmbarc", tipo: tCatCFDI_cPais},
		{nombre: "NombreEmbarcador", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}DerechosDePaso
var eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "DerechosDePaso", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDerechoDePaso", tipo: tCatCartaPorte_cDerechosDePaso, requerido: true},
		{nombre: "KilometrajePagado", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Contenedor
var eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedor, requerido: true},
		{nombre: "PesoContenedorVacio", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio, requerido: true},
		{nombre: "PesoNetoMercancia", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Carro
var eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Carro", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoCarro", tipo: tCatCartaPorte_cTipoCarro, requerido: true},
		{nombre: "MatriculaCarro", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro, requerido: true},
		{nombre: "GuiaCarro", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro, requerido: true},
		{nombre: "ToneladasNetasCarro", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor, min: 0, max: -1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}TransporteFerroviario
var eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "TransporteFerroviario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoDeServicio", tipo: tCatCartaPorte_cTipoDeServicio, requerido: true},
		{nombre: "TipoDeTrafico", tipo: tCatCartaPorte_cTipoDeTrafico, requerido: true},
		{nombre: "NombreAseg", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso, min: 0, max: -1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario_Carro, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Mercancias
var eCartaPorte30_CartaPorte_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PesoBrutoTotal", tipo: tCartaPorte30_CartaPorte_Mercancias_PesoBrutoTotal, requerido: true},
		{nombre: "UnidadPeso", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoNetoTotal", tipo: tCartaPorte30_CartaPorte_Mercancias_PesoNetoTotal},
		{nombre: "NumTotalMercancias", tipo: tCartaPorte30_CartaPorte_Mercancias_NumTotalMercancias, requerido: true},
		{nombre: "CargoPorTasacion", tipo: tTdCFDI_tImporte},
		{nombre: "LogisticaInversaRecoleccionDevolucion", tipo: tCartaPorte30_CartaPorte_Mercancias_LogisticaInversaRecoleccionDevolucion},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Mercancias_Mercancia, min: 1, max: -1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_Autotransporte, min: 0, max: 1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteMaritimo, min: 0, max: 1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteAereo, min: 0, max: 1},
		{elem: eCartaPorte30_CartaPorte_Mercancias_TransporteFerroviario, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}PartesTransporte
var eCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_PartesTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "PartesTransporte", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ParteTransporte", tipo: tCatCartaPorte_cParteTransporte, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}Domicilio
var eCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}TiposFigura
var eCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "TiposFigura", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoFigura", tipo: tCatCartaPorte_cFiguraTransporte, requerido: true},
		{nombre: "RFCFigura", tipo: tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_RFCFigura},
		{nombre: "NumLicencia", tipo: tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_NumLicencia},
		{nombre: "NombreFigura", tipo: tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_NombreFigura, requerido: true},
		{nombre: "NumRegIdTribFigura", tipo: tCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_NumRegIdTribFigura},
		{nombre: "ResidenciaFiscalFigura", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_PartesTransporte, min: 0, max: -1},
		{elem: eCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}FiguraTransporte
var eCartaPorte30_CartaPorte_FiguraTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "FiguraTransporte", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_FiguraTransporte_TiposFigura, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte30.xsd: {http://www.sat.gob.mx/CartaPorte30}CartaPorte
var eCartaPorte30_CartaPorte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte30", nombre: "CartaPorte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "3.0", conFijo: true},
		{nombre: "IdCCP", tipo: tCartaPorte30_CartaPorte_IdCCP, requerido: true},
		{nombre: "TranspInternac", tipo: tCartaPorte30_CartaPorte_TranspInternac, requerido: true},
		{nombre: "RegimenAduanero", tipo: tCatCartaPorte_cRegimenAduanero},
		{nombre: "EntradaSalidaMerc", tipo: tCartaPorte30_CartaPorte_EntradaSalidaMerc},
		{nombre: "PaisOrigenDestino", tipo: tCatCFDI_cPais},
		{nombre: "ViaEntradaSalida", tipo: tCatCartaPorte_cCveTransporte},
		{nombre: "TotalDistRec", tipo: tCartaPorte30_CartaPorte_TotalDistRec},
		{nombre: "RegistroISTMO", tipo: tCartaPorte30_CartaPorte_RegistroISTMO},
		{nombre: "UbicacionPoloOrigen", tipo: tCatCartaPorte_cRegistroISTMO},
		{nombre: "UbicacionPoloDestino", tipo: tCatCartaPorte_cRegistroISTMO},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte30_CartaPorte_Ubicaciones, min: 1, max: 1},
		{elem: eCartaPorte30_CartaPorte_Mercancias, min: 1, max: 1},
		{elem: eCartaPorte30_CartaPorte_FiguraTransporte, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Domicilio
var eCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Ubicacion
var eCartaPorte20_CartaPorte_Ubicaciones_Ubicacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Ubicacion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoUbicacion", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_TipoUbicacion, requerido: true},
		{nombre: "IDUbicacion", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_IDUbicacion},
		{nombre: "RFCRemitenteDestinatario", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "NombreRemitenteDestinatario", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NombreRemitenteDestinatario},
		{nombre: "NumRegIdTrib", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NumRegIdTrib},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumEstacion", tipo: tCatCartaPorte_cEstaciones},
		{nombre: "NombreEstacion", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NombreEstacion},
		{nombre: "NavegacionTrafico", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_NavegacionTrafico},
		{nombre: "FechaHoraSalidaLlegada", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "TipoEstacion", tipo: tCatCartaPorte_cTipoEstacion},
		{nombre: "DistanciaRecorrida", tipo: tCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Ubicaciones_Ubicacion_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Ubicaciones
var eCartaPorte20_CartaPorte_Ubicaciones = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Ubicaciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Ubicaciones_Ubicacion, min: 2, max: -1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Pedimentos
var eCartaPorte20_CartaPorte_Mercancias_Mercancia_Pedimentos = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Pedimentos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Pedimento", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_Pedimentos_Pedimento, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}GuiasIdentificacion
var eCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "GuiasIdentificacion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumeroGuiaIdentificacion", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_NumeroGuiaIdentificacion, requerido: true},
		{nombre: "DescripGuiaIdentificacion", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_DescripGuiaIdentificacion, requerido: true},
		{nombre: "PesoGuiaIdentificacion", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion_PesoGuiaIdentificacion, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}CantidadTransporta
var eCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "CantidadTransporta", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Cantidad", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad, requerido: true},
		{nombre: "IDOrigen", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen, requerido: true},
		{nombre: "IDDestino", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino, requerido: true},
		{nombre: "CvesTransporte", tipo: tCatCartaPorte_cCveTransporte},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}DetalleMercancia
var eCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "DetalleMercancia", vacio: true,
	attrs: []atributoXSD{
		{nombre: "UnidadPesoMerc", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoBruto", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto, requerido: true},
		{nombre: "PesoNeto", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto, requerido: true},
		{nombre: "PesoTara", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara, requerido: true},
		{nombre: "NumPiezas", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Mercancia
var eCartaPorte20_CartaPorte_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "BienesTransp", tipo: tCatCartaPorte_cClaveProdServCP, requerido: true},
		{nombre: "ClaveSTCC", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_ClaveSTCC},
		{nombre: "Descripcion", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_Descripcion, requerido: true},
		{nombre: "Cantidad", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_Cantidad, requerido: true},
		{nombre: "ClaveUnidad", tipo: tCatCFDI_cClaveUnidad, requerido: true},
		{nombre: "Unidad", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_Unidad},
		{nombre: "Dimensiones", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_Dimensiones},
		{nombre: "MaterialPeligroso", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
		{nombre: "CveMaterialPeligroso", tipo: tCatCartaPorte_cMaterialPeligroso},
		{nombre: "Embalaje", tipo: tCatCartaPorte_cTipoEmbalaje},
		{nombre: "DescripEmbalaje", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_DescripEmbalaje},
		{nombre: "PesoEnKg", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_PesoEnKg, requerido: true},
		{nombre: "ValorMercancia", tipo: tTdCFDI_tImporte},
		{nombre: "Moneda", tipo: tCatCFDI_cMoneda},
		{nombre: "FraccionArancelaria", tipo: tCatComExt_cFraccionArancelaria},
		{nombre: "UUIDComercioExt", tipo: tCartaPorte20_CartaPorte_Mercancias_Mercancia_UUIDComercioExt},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_Mercancia_Pedimentos, min: 0, max: -1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_Mercancia_GuiasIdentificacion, min: 0, max: -1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_Mercancia_CantidadTransporta, min: 0, max: -1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_Mercancia_DetalleMercancia, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}IdentificacionVehicular
var eCartaPorte20_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "IdentificacionVehicular", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ConfigVehicular", tipo: tCatCartaPorte_cConfigAutotransporte, requerido: true},
		{nombre: "PlacaVM", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_PlacaVM, requerido: true},
		{nombre: "AnioModeloVM", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular_AnioModeloVM, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Seguros
var eCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Seguros", vacio: true,
	attrs: []atributoXSD{
		{nombre: "AseguraRespCivil", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraRespCivil, requerido: true},
		{nombre: "PolizaRespCivil", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaRespCivil, requerido: true},
		{nombre: "AseguraMedAmbiente", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraMedAmbiente},
		{nombre: "PolizaMedAmbiente", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaMedAmbiente},
		{nombre: "AseguraCarga", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_AseguraCarga},
		{nombre: "PolizaCarga", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros_PolizaCarga},
		{nombre: "PrimaSeguro", tipo: tTdCFDI_tImporte},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Remolque
var eCartaPorte20_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Remolque", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubTipoRem", tipo: tCatCartaPorte_cSubTipoRem, requerido: true},
		{nombre: "Placa", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque_Placa, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Remolques
var eCartaPorte20_CartaPorte_Mercancias_Autotransporte_Remolques = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Remolques", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_Autotransporte_Remolques_Remolque, min: 1, max: 2},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Autotransporte
var eCartaPorte20_CartaPorte_Mercancias_Autotransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Autotransporte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte20_CartaPorte_Mercancias_Autotransporte_NumPermisoSCT, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_Autotransporte_IdentificacionVehicular, min: 1, max: 1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_Autotransporte_Seguros, min: 1, max: 1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_Autotransporte_Remolques, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Contenedor
var eCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "MatriculaContenedor", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor, requerido: true},
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedorMaritimo, requerido: true},
		{nombre: "NumPrecinto", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}TransporteMaritimo
var eCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "TransporteMaritimo", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT},
		{nombre: "NombreAseg", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro},
		{nombre: "TipoEmbarcacion", tipo: tCatCartaPorte_cConfigMaritima, requerido: true},
		{nombre: "Matricula", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Matricula, requerido: true},
		{nombre: "NumeroOMI", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI, requerido: true},
		{nombre: "AnioEmbarcacion", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion},
		{nombre: "NombreEmbarc", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc},
		{nombre: "NacionalidadEmbarc", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "UnidadesDeArqBruto", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto, requerido: true},
		{nombre: "TipoCarga", tipo: tCatCartaPorte_cClaveTipoCarga, requerido: true},
		{nombre: "NumCertITC", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumCertITC, requerido: true},
		{nombre: "Eslora", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Eslora},
		{nombre: "Manga", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Manga},
		{nombre: "Calado", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Calado},
		{nombre: "LineaNaviera", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera},
		{nombre: "NombreAgenteNaviero", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero, requerido: true},
		{nombre: "NumAutorizacionNaviero", tipo: tCatCartaPorte_cNumAutorizacionNaviero, requerido: true},
		{nombre: "NumViaje", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumViaje},
		{nombre: "NumConocEmbarc", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo_Contenedor, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}TransporteAereo
var eCartaPorte20_CartaPorte_Mercancias_TransporteAereo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "TransporteAereo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT, requerido: true},
		{nombre: "MatriculaAeronave", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave},
		{nombre: "NombreAseg", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro},
		{nombre: "NumeroGuia", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumeroGuia, requerido: true},
		{nombre: "LugarContrato", tipo: tTdCFDI_tDescrip120},
		{nombre: "CodigoTransportista", tipo: tCatCartaPorte_cCodigoTransporteAereo, requerido: true},
		{nombre: "RFCEmbarcador", tipo: tTdCFDI_tRFC},
		{nombre: "NumRegIdTribEmbarc", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc},
		{nombre: "ResidenciaFiscalEmbarc", tipo: tCatCFDI_cPais},
		{nombre: "NombreEmbarcador", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}DerechosDePaso
var eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "DerechosDePaso", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDerechoDePaso", tipo: tCatCartaPorte_cDerechosDePaso, requerido: true},
		{nombre: "KilometrajePagado", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Contenedor
var eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedor, requerido: true},
		{nombre: "PesoContenedorVacio", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio, requerido: true},
		{nombre: "PesoNetoMercancia", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Carro
var eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Carro", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoCarro", tipo: tCatCartaPorte_cTipoCarro, requerido: true},
		{nombre: "MatriculaCarro", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro, requerido: true},
		{nombre: "GuiaCarro", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro, requerido: true},
		{nombre: "ToneladasNetasCarro", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor, min: 0, max: -1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}TransporteFerroviario
var eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "TransporteFerroviario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoDeServicio", tipo: tCatCartaPorte_cTipoDeServicio, requerido: true},
		{nombre: "TipoDeTrafico", tipo: tCatCartaPorte_cTipoDeTrafico, requerido: true},
		{nombre: "NombreAseg", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso, min: 0, max: -1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario_Carro, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Mercancias
var eCartaPorte20_CartaPorte_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PesoBrutoTotal", tipo: tCartaPorte20_CartaPorte_Mercancias_PesoBrutoTotal, requerido: true},
		{nombre: "UnidadPeso", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoNetoTotal", tipo: tCartaPorte20_CartaPorte_Mercancias_PesoNetoTotal},
		{nombre: "NumTotalMercancias", tipo: tCartaPorte20_CartaPorte_Mercancias_NumTotalMercancias, requerido: true},
		{nombre: "CargoPorTasacion", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Mercancias_Mercancia, min: 1, max: -1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_Autotransporte, min: 0, max: 1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteMaritimo, min: 0, max: 1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteAereo, min: 0, max: 1},
		{elem: eCartaPorte20_CartaPorte_Mercancias_TransporteFerroviario, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}PartesTransporte
var eCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_PartesTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "PartesTransporte", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ParteTransporte", tipo: tCatCartaPorte_cParteTransporte, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}Domicilio
var eCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}TiposFigura
var eCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "TiposFigura", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoFigura", tipo: tCatCartaPorte_cFiguraTransporte, requerido: true},
		{nombre: "RFCFigura", tipo: tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_RFCFigura},
		{nombre: "NumLicencia", tipo: tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_NumLicencia},
		{nombre: "NombreFigura", tipo: tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_NombreFigura},
		{nombre: "NumRegIdTribFigura", tipo: tCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_NumRegIdTribFigura},
		{nombre: "ResidenciaFiscalFigura", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_PartesTransporte, min: 0, max: -1},
		{elem: eCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}FiguraTransporte
var eCartaPorte20_CartaPorte_FiguraTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "FiguraTransporte", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_FiguraTransporte_TiposFigura, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte20.xsd: {http://www.sat.gob.mx/CartaPorte20}CartaPorte
var eCartaPorte20_CartaPorte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte20", nombre: "CartaPorte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "2.0", conFijo: true},
		{nombre: "TranspInternac", tipo: tCartaPorte20_CartaPorte_TranspInternac, requerido: true},
		{nombre: "EntradaSalidaMerc", tipo: tCartaPorte20_CartaPorte_EntradaSalidaMerc},
		{nombre: "PaisOrigenDestino", tipo: tCatCFDI_cPais},
		{nombre: "ViaEntradaSalida", tipo: tCatCartaPorte_cCveTransporte},
		{nombre: "TotalDistRec", tipo: tCartaPorte20_CartaPorte_TotalDistRec},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte20_CartaPorte_Ubicaciones, min: 1, max: 1},
		{elem: eCartaPorte20_CartaPorte_Mercancias, min: 1, max: 1},
		{elem: eCartaPorte20_CartaPorte_FiguraTransporte, min: 0, max: 1},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Domicilio
var eComExt20_ComercioExterior_Emisor_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tCatCFDI_cColonia},
		{nombre: "Localidad", tipo: tCatCFDI_cLocalidad},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tCatCFDI_cMunicipio},
		{nombre: "Estado", tipo: tCatCFDI_cEstado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCatCFDI_cCodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Emisor
var eComExt20_ComercioExterior_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Emisor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Curp", tipo: tTdCFDI_tCURP},
	},
	particulas: []particulaXSD{
		{elem: eComExt20_ComercioExterior_Emisor_Domicilio, min: 1, max: 1},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Propietario
var eComExt20_ComercioExterior_Propietario = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Propietario", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt20_ComercioExterior_Propietario_NumRegIdTrib, requerido: true},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais, requerido: true},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Domicilio
var eComExt20_ComercioExterior_Receptor_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tComExt20_ComercioExterior_Receptor_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tComExt20_ComercioExterior_Receptor_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Receptor
var eComExt20_ComercioExterior_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Receptor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt20_ComercioExterior_Receptor_NumRegIdTrib},
	},
	particulas: []particulaXSD{
		{elem: eComExt20_ComercioExterior_Receptor_Domicilio, min: 0, max: 1},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Domicilio
var eComExt20_ComercioExterior_Destinatario_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tComExt20_ComercioExterior_Destinatario_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tComExt20_ComercioExterior_Destinatario_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Destinatario
var eComExt20_ComercioExterior_Destinatario = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Destinatario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt20_ComercioExterior_Destinatario_NumRegIdTrib},
		{nombre: "Nombre", tipo: tComExt20_ComercioExterior_Destinatario_Nombre},
	},
	particulas: []particulaXSD{
		{elem: eComExt20_ComercioExterior_Destinatario_Domicilio, min: 1, max: -1},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}DescripcionesEspecificas
var eComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "DescripcionesEspecificas", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Marca", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Marca, requerido: true},
		{nombre: "Modelo", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Modelo},
		{nombre: "SubModelo", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_SubModelo},
		{nombre: "NumeroSerie", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_NumeroSerie},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Mercancia
var eComExt20_ComercioExterior_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NoIdentificacion", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_NoIdentificacion, requerido: true},
		{nombre: "FraccionArancelaria", tipo: tCatComExt_cFraccionArancelaria},
		{nombre: "CantidadAduana", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_CantidadAduana},
		{nombre: "UnidadAduana", tipo: tCatComExt_cUnidadAduana},
		{nombre: "ValorUnitarioAduana", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_ValorUnitarioAduana},
		{nombre: "ValorDolares", tipo: tComExt20_ComercioExterior_Mercancias_Mercancia_ValorDolares, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eComExt20_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas, min: 0, max: -1},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}Mercancias
var eComExt20_ComercioExterior_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eComExt20_ComercioExterior_Mercancias_Mercancia, min: 1, max: -1},
	}}

// cfd/ComercioExterior20/ComercioExterior20.xsd: {http://www.sat.gob.mx/ComercioExterior20}ComercioExterior
var eComExt20_ComercioExterior = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior20", nombre: "ComercioExterior", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tComExt20_ComercioExterior_Version, requerido: true, fijo: "2.0", conFijo: true},
		{nombre: "MotivoTraslado", tipo: tCatComExt_cMotivoTraslado},
		{nombre: "ClaveDePedimento", tipo: tCatComExt_cClavePedimento, requerido: true},
		{nombre: "CertificadoOrigen", tipo: tComExt20_ComercioExterior_CertificadoOrigen, requerido: true},
		{nombre: "NumCertificadoOrigen", tipo: tComExt20_ComercioExterior_NumCertificadoOrigen},
		{nombre: "NumeroExportadorConfiable", tipo: tComExt20_ComercioExterior_NumeroExportadorConfiable},
		{nombre: "Incoterm", tipo: tCatComExt_cINCOTERM},
		{nombre: "Observaciones", tipo: tComExt20_ComercioExterior_Observaciones},
		{nombre: "TipoCambioUSD", tipo: tTdCFDI_tTipoCambio, requerido: true},
		{nombre: "TotalUSD", tipo: tTdCFDI_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eComExt20_ComercioExterior_Emisor, min: 0, max: 1},
		{elem: eComExt20_ComercioExterior_Propietario, min: 0, max: -1},
		{elem: eComExt20_ComercioExterior_Receptor, min: 0, max: 1},
		{elem: eComExt20_ComercioExterior_Destinatario, min: 0, max: -1},
		{elem: eComExt20_ComercioExterior_Mercancias, min: 1, max: 1},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Domicilio
var eComExt11_ComercioExterior_Emisor_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tCatCFDI_cColonia},
		{nombre: "Localidad", tipo: tCatCFDI_cLocalidad},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tCatCFDI_cMunicipio},
		{nombre: "Estado", tipo: tCatCFDI_cEstado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCatCFDI_cCodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Emisor
var eComExt11_ComercioExterior_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Emisor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Curp", tipo: tTdCFDI_tCURP},
	},
	particulas: []particulaXSD{
		{elem: eComExt11_ComercioExterior_Emisor_Domicilio, min: 0, max: 1},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Propietario
var eComExt11_ComercioExterior_Propietario = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Propietario", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt11_ComercioExterior_Propietario_NumRegIdTrib, requerido: true},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais, requerido: true},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Domicilio
var eComExt11_ComercioExterior_Receptor_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tComExt11_ComercioExterior_Receptor_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tComExt11_ComercioExterior_Receptor_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Receptor
var eComExt11_ComercioExterior_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Receptor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt11_ComercioExterior_Receptor_NumRegIdTrib},
	},
	particulas: []particulaXSD{
		{elem: eComExt11_ComercioExterior_Receptor_Domicilio, min: 0, max: 1},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Domicilio
var eComExt11_ComercioExterior_Destinatario_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tComExt11_ComercioExterior_Destinatario_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tComExt11_ComercioExterior_Destinatario_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Destinatario
var eComExt11_ComercioExterior_Destinatario = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Destinatario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt11_ComercioExterior_Destinatario_NumRegIdTrib},
		{nombre: "Nombre", tipo: tComExt11_ComercioExterior_Destinatario_Nombre},
	},
	particulas: []particulaXSD{
		{elem: eComExt11_ComercioExterior_Destinatario_Domicilio, min: 1, max: -1},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}DescripcionesEspecificas
var eComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "DescripcionesEspecificas", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Marca", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Marca, requerido: true},
		{nombre: "Modelo", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Modelo},
		{nombre: "SubModelo", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_SubModelo},
		{nombre: "NumeroSerie", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_NumeroSerie},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Mercancia
var eComExt11_ComercioExterior_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NoIdentificacion", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_NoIdentificacion, requerido: true},
		{nombre: "FraccionArancelaria", tipo: tCatComExt_cFraccionArancelaria},
		{nombre: "CantidadAduana", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_CantidadAduana},
		{nombre: "UnidadAduana", tipo: tCatComExt_cUnidadAduana},
		{nombre: "ValorUnitarioAduana", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_ValorUnitarioAduana},
		{nombre: "ValorDolares", tipo: tComExt11_ComercioExterior_Mercancias_Mercancia_ValorDolares, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eComExt11_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas, min: 0, max: -1},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}Mercancias
var eComExt11_ComercioExterior_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eComExt11_ComercioExterior_Mercancias_Mercancia, min: 1, max: -1},
	}}

// cfd/ComercioExterior11/ComercioExterior11.xsd: {http://www.sat.gob.mx/ComercioExterior11}ComercioExterior
var eComExt11_ComercioExterior = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior11", nombre: "ComercioExterior", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tComExt11_ComercioExterior_Version, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "MotivoTraslado", tipo: tCatComExt_cMotivoTraslado},
		{nombre: "TipoOperacion", tipo: tCatComExt_cTipoOperacion, requerido: true},
		{nombre: "ClaveDePedimento", tipo: tCatComExt_cClavePedimento},
		{nombre: "CertificadoOrigen", tipo: tComExt11_ComercioExterior_CertificadoOrigen},
		{nombre: "NumCertificadoOrigen", tipo: tComExt11_ComercioExterior_NumCertificadoOrigen},
		{nombre: "NumeroExportadorConfiable", tipo: tComExt11_ComercioExterior_NumeroExportadorConfiable},
		{nombre: "Incoterm", tipo: tCatComExt_cINCOTERM},
		{nombre: "Subdivision", tipo: tComExt11_ComercioExterior_Subdivision},
		{nombre: "Observaciones", tipo: tComExt11_ComercioExterior_Observaciones},
		{nombre: "TipoCambioUSD", tipo: tTdCFDI_tTipoCambio},
		{nombre: "TotalUSD", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eComExt11_ComercioExterior_Emisor, min: 0, max: 1},
		{elem: eComExt11_ComercioExterior_Propietario, min: 0, max: -1},
		{elem: eComExt11_ComercioExterior_Receptor, min: 0, max: 1},
		{elem: eComExt11_ComercioExterior_Destinatario, min: 0, max: -1},
		{elem: eComExt11_ComercioExterior_Mercancias, min: 0, max: 1},
	}}

// cfd/hidrocarburospetroliferos.xsd: {http://www.sat.gob.mx/hidrocarburospetroliferos}HidroYPetro
var eHidroYPetro_HidroYPetro = &elementoXSD{ns: "http://www.sat.gob.mx/hidrocarburospetroliferos", nombre: "HidroYPetro", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TipoPermiso", tipo: tCatHidroYPetro_cTipoPermiso, requerido: true},
		{nombre: "NumeroPermiso", tipo: tHidroYPetro_HidroYPetro_NumeroPermiso, requerido: true},
		{nombre: "ClaveHYP", tipo: tCatHidroYPetro_cClaveHYP, requerido: true},
		{nombre: "SubProductoHYP", tipo: tCatHidroYPetro_cSubProductoHYP, requerido: true},
	}}

// cfd/cfdiregistrofiscal/cfdiregistrofiscal.xsd: {http://www.sat.gob.mx/registrofiscal}CFDIRegistroFiscal
var eRegistroFiscal_CFDIRegistroFiscal = &elementoXSD{ns: "http://www.sat.gob.mx/registrofiscal", nombre: "CFDIRegistroFiscal", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "Folio", tipo: tRegistroFiscal_CFDIRegistroFiscal_Folio, requerido: true},
	}}

// cfd/divisas/divisas.xsd: {http://www.sat.gob.mx/divisas}Divisas
var eDivisas_Divisas = &elementoXSD{ns: "http://www.sat.gob.mx/divisas", nombre: "Divisas", vacio: true,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "tipoOperacion", tipo: tDivisas_Divisas_tipoOperacion, requerido: true},
	}}

// cfd/EstadoDeCuentaCombustible/ecc12.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible12}Traslado
var eECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible12", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Impuesto, requerido: true},
		{nombre: "TasaOCuota", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_TasaOCuota, requerido: true},
		{nombre: "Importe", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Importe, requerido: true},
	}}

// cfd/EstadoDeCuentaCombustible/ecc12.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible12}Traslados
var eECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible12", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc12.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible12}ConceptoEstadoDeCuentaCombustible
var eECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible12", nombre: "ConceptoEstadoDeCuentaCombustible", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Identificador", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Identificador, requerido: true},
		{nombre: "Fecha", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "Rfc", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "ClaveEstacion", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ClaveEstacion, requerido: true},
		{nombre: "Cantidad", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Cantidad, requerido: true},
		{nombre: "TipoCombustible", tipo: tCatCombustible_cClaveTipoCombustible, requerido: true},
		{nombre: "Unidad", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Unidad},
		{nombre: "NombreCombustible", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_NombreCombustible, requerido: true},
		{nombre: "FolioOperacion", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_FolioOperacion, requerido: true},
		{nombre: "ValorUnitario", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ValorUnitario, requerido: true},
		{nombre: "Importe", tipo: tECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Importe, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados, min: 1, max: 1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc12.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible12}Conceptos
var eECC12_EstadoDeCuentaCombustible_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible12", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eECC12_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible, min: 1, max: -1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc12.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible12}EstadoDeCuentaCombustible
var eECC12_EstadoDeCuentaCombustible = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible12", nombre: "EstadoDeCuentaCombustible", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tECC12_EstadoDeCuentaCombustible_Version, requerido: true, fijo: "1.2", conFijo: true},
		{nombre: "TipoOperacion", tipo: tECC12_EstadoDeCuentaCombustible_TipoOperacion, requerido: true, fijo: "Tarjeta", conFijo: true},
		{nombre: "NumeroDeCuenta", tipo: tECC12_EstadoDeCuentaCombustible_NumeroDeCuenta, requerido: true},
		{nombre: "SubTotal", tipo: tECC12_EstadoDeCuentaCombustible_SubTotal, requerido: true},
		{nombre: "Total", tipo: tECC12_EstadoDeCuentaCombustible_Total, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eECC12_EstadoDeCuentaCombustible_Conceptos, min: 1, max: 1},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}InformacionFiscalTercero
var eTerceros11_PorCuentadeTerceros_InformacionFiscalTercero = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "InformacionFiscalTercero", vacio: true,
	attrs: []atributoXSD{
		{nombre: "calle", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_calle, requerido: true},
		{nombre: "noExterior", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_noExterior},
		{nombre: "noInterior", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_noInterior},
		{nombre: "colonia", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_colonia},
		{nombre: "localidad", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_localidad},
		{nombre: "referencia", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_referencia},
		{nombre: "municipio", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_municipio, requerido: true},
		{nombre: "estado", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_estado, requerido: true},
		{nombre: "pais", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_pais, requerido: true},
		{nombre: "codigoPostal", tipo: tTerceros11_PorCuentadeTerceros_InformacionFiscalTercero_codigoPostal, requerido: true},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}InformacionAduanera
var eTerceros11_PorCuentadeTerceros_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tTerceros11_PorCuentadeTerceros_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tTerceros11_PorCuentadeTerceros_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tTerceros11_PorCuentadeTerceros_InformacionAduanera_aduana},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}InformacionAduanera
var eTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera_aduana},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}Parte
var eTerceros11_PorCuentadeTerceros_Parte = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "Parte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "cantidad", tipo: tTerceros11_PorCuentadeTerceros_Parte_cantidad, requerido: true},
		{nombre: "unidad", tipo: tTerceros11_PorCuentadeTerceros_Parte_unidad},
		{nombre: "noIdentificacion", tipo: tTerceros11_PorCuentadeTerceros_Parte_noIdentificacion},
		{nombre: "descripcion", tipo: tTerceros11_PorCuentadeTerceros_Parte_descripcion, requerido: true},
		{nombre: "valorUnitario", tipo: tTerceros11_tImporte},
		{nombre: "importe", tipo: tTerceros11_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera, min: 0, max: -1},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}CuentaPredial
var eTerceros11_PorCuentadeTerceros_CuentaPredial = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "CuentaPredial", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tTerceros11_PorCuentadeTerceros_CuentaPredial_numero, requerido: true},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}Retencion
var eTerceros11_PorCuentadeTerceros_Impuestos_Retenciones_Retencion = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "Retencion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "impuesto", tipo: tTerceros11_PorCuentadeTerceros_Impuestos_Retenciones_Retencion_impuesto, requerido: true},
		{nombre: "importe", tipo: tTerceros11_tImporte, requerido: true},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}Retenciones
var eTerceros11_PorCuentadeTerceros_Impuestos_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eTerceros11_PorCuentadeTerceros_Impuestos_Retenciones_Retencion, min: 1, max: -1},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}Traslado
var eTerceros11_PorCuentadeTerceros_Impuestos_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "impuesto", tipo: tTerceros11_PorCuentadeTerceros_Impuestos_Traslados_Traslado_impuesto, requerido: true},
		{nombre: "tasa", tipo: tTerceros11_tImporte, requerido: true},
		{nombre: "importe", tipo: tTerceros11_tImporte, requerido: true},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}Traslados
var eTerceros11_PorCuentadeTerceros_Impuestos_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eTerceros11_PorCuentadeTerceros_Impuestos_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}Impuestos
var eTerceros11_PorCuentadeTerceros_Impuestos = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "Impuestos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eTerceros11_PorCuentadeTerceros_Impuestos_Retenciones, min: 0, max: 1},
		{elem: eTerceros11_PorCuentadeTerceros_Impuestos_Traslados, min: 0, max: 1},
	}}

// cfd/terceros/terceros11.xsd: {http://www.sat.gob.mx/terceros}PorCuentadeTerceros
var eTerceros11_PorCuentadeTerceros = &elementoXSD{ns: "http://www.sat.gob.mx/terceros", nombre: "PorCuentadeTerceros", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsanySimpleType, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "rfc", tipo: tTerceros11_tRFC, requerido: true},
		{nombre: "nombre", tipo: tTerceros11_PorCuentadeTerceros_nombre},
	},
	particulas: []particulaXSD{
		{elem: eTerceros11_PorCuentadeTerceros_InformacionFiscalTercero, min: 0, max: 1},
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eTerceros11_PorCuentadeTerceros_InformacionAduanera, min: 0, max: 1},
			{elem: eTerceros11_PorCuentadeTerceros_Parte, min: 0, max: -1},
			{elem: eTerceros11_PorCuentadeTerceros_CuentaPredial, min: 0, max: 1},
		}},
		{elem: eTerceros11_PorCuentadeTerceros_Impuestos, min: 1, max: 1},
	}}

// cfd/ventavehiculos/ventavehiculos11.xsd: {http://www.sat.gob.mx/ventavehiculos}InformacionAduanera
var eVentaVehiculos11_VentaVehiculos_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tVentaVehiculos11_VentaVehiculos_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tVentaVehiculos11_VentaVehiculos_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tVentaVehiculos11_VentaVehiculos_InformacionAduanera_aduana},
	}}

// cfd/ventavehiculos/ventavehiculos11.xsd: {http://www.sat.gob.mx/ventavehiculos}InformacionAduanera
var eVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera_aduana},
	}}

// cfd/ventavehiculos/ventavehiculos11.xsd: {http://www.sat.gob.mx/ventavehiculos}Parte
var eVentaVehiculos11_VentaVehiculos_Parte = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "Parte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "cantidad", tipo: tVentaVehiculos11_VentaVehiculos_Parte_cantidad, requerido: true},
		{nombre: "unidad", tipo: tVentaVehiculos11_VentaVehiculos_Parte_unidad},
		{nombre: "noIdentificacion", tipo: tVentaVehiculos11_VentaVehiculos_Parte_noIdentificacion},
		{nombre: "descripcion", tipo: tVentaVehiculos11_VentaVehiculos_Parte_descripcion, requerido: true},
		{nombre: "valorUnitario", tipo: tVentaVehiculos11_tImporte},
		{nombre: "importe", tipo: tVentaVehiculos11_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eVentaVehiculos11_VentaVehiculos_Parte_InformacionAduanera, min: 0, max: -1},
	}}

// cfd/ventavehiculos/ventavehiculos11.xsd: {http://www.sat.gob.mx/ventavehiculos}VentaVehiculos
var eVentaVehiculos11_VentaVehiculos = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "VentaVehiculos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsanySimpleType, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "ClaveVehicular", tipo: tVentaVehiculos11_VentaVehiculos_ClaveVehicular, requerido: true},
		{nombre: "Niv", tipo: tVentaVehiculos11_VentaVehiculos_Niv, requerido: true},
	},
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eVentaVehiculos11_VentaVehiculos_InformacionAduanera, min: 0, max: -1},
			{elem: eVentaVehiculos11_VentaVehiculos_Parte, min: 0, max: -1},
		}},
	}}

// cfd/aerolineas/aerolineas.xsd: {http://www.sat.gob.mx/aerolineas}Cargo
var eAerolineas_Aerolineas_OtrosCargos_Cargo = &elementoXSD{ns: "http://www.sat.gob.mx/aerolineas", nombre: "Cargo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "CodigoCargo", tipo: tAerolineas_Aerolineas_OtrosCargos_Cargo_CodigoCargo, requerido: true},
		{nombre: "Importe", tipo: tAerolineas_tImporte, requerido: true},
	}}

// cfd/aerolineas/aerolineas.xsd: {http://www.sat.gob.mx/aerolineas}OtrosCargos
var eAerolineas_Aerolineas_OtrosCargos = &elementoXSD{ns: "http://www.sat.gob.mx/aerolineas", nombre: "OtrosCargos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalCargos", tipo: tAerolineas_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eAerolineas_Aerolineas_OtrosCargos_Cargo, min: 1, max: -1},
	}}

// cfd/aerolineas/aerolineas.xsd: {http://www.sat.gob.mx/aerolineas}Aerolineas
var eAerolineas_Aerolineas = &elementoXSD{ns: "http://www.sat.gob.mx/aerolineas", nombre: "Aerolineas", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TUA", tipo: tAerolineas_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eAerolineas_Aerolineas_OtrosCargos, min: 0, max: 1},
	}}

// cfd/leyendasFiscales/leyendasFisc.xsd: {http://www.sat.gob.mx/leyendasFiscales}Leyenda
var eLeyendasFisc_LeyendasFiscales_Leyenda = &elementoXSD{ns: "http://www.sat.gob.mx/leyendasFiscales", nombre: "Leyenda", vacio: true,
	attrs: []atributoXSD{
		{nombre: "disposicionFiscal", tipo: tLeyendasFisc_LeyendasFiscales_Leyenda_disposicionFiscal},
		{nombre: "norma", tipo: tLeyendasFisc_LeyendasFiscales_Leyenda_norma},
		{nombre: "textoLeyenda", tipo: tLeyendasFisc_LeyendasFiscales_Leyenda_textoLeyenda, requerido: true},
	}}

// cfd/leyendasFiscales/leyendasFisc.xsd: {http://www.sat.gob.mx/leyendasFiscales}LeyendasFiscales
var eLeyendasFisc_LeyendasFiscales = &elementoXSD{ns: "http://www.sat.gob.mx/leyendasFiscales", nombre: "LeyendasFiscales", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
	},
	particulas: []particulaXSD{
		{elem: eLeyendasFisc_LeyendasFiscales_Leyenda, min: 1, max: -1},
	}}

// cfd/iedu/iedu.xsd: {http://www.sat.gob.mx/iedu}instEducativas
var eIEDU_instEducativas = &elementoXSD{ns: "http://www.sat.gob.mx/iedu", nombre: "instEducativas", vacio: true,
	attrs: []atributoXSD{
		{nombre: "version", tipo: tIEDU_instEducativas_version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "nombreAlumno", tipo: tIEDU_instEducativas_nombreAlumno, requerido: true},
		{nombre: "CURP", tipo: tIEDU_tCURP, requerido: true},
		{nombre: "nivelEducativo", tipo: tIEDU_instEducativas_nivelEducativo, requerido: true},
		{nombre: "autRVOE", tipo: tIEDU_instEducativas_autRVOE, requerido: true},
		{nombre: "rfcPago", tipo: tIEDU_tRFC},
	}}

// cfd/donat/donat11.xsd: {http://www.sat.gob.mx/donat}Donatarias
var eDonat11_Donatarias = &elementoXSD{ns: "http://www.sat.gob.mx/donat", nombre: "Donatarias", vacio: true,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "noAutorizacion", tipo: xsstring, requerido: true},
		{nombre: "fechaAutorizacion", tipo: xsdate, requerido: true},
		{nombre: "leyenda", tipo: tDonat11_Donatarias_leyenda, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Origen
var eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Origen", vacio: true,
	attrs: []atributoXSD{
		{nombre: "IDOrigen", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_IDOrigen},
		{nombre: "RFCRemitente", tipo: tTdCFDI_tRFC},
		{nombre: "NombreRemitente", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NombreRemitente},
		{nombre: "NumRegIdTrib", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NumRegIdTrib},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumEstacion", tipo: tCatCartaPorte_cEstaciones},
		{nombre: "NombreEstacion", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NombreEstacion},
		{nombre: "NavegacionTrafico", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen_NavegacionTrafico},
		{nombre: "FechaHoraSalida", tipo: tTdCFDI_tFechaH, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Destino
var eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Destino", vacio: true,
	attrs: []atributoXSD{
		{nombre: "IDDestino", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_IDDestino},
		{nombre: "RFCDestinatario", tipo: tTdCFDI_tRFC},
		{nombre: "NombreDestinatario", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NombreDestinatario},
		{nombre: "NumRegIdTrib", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NumRegIdTrib},
		{nombre: "ResidenciaFiscal", tipo: tCatCFDI_cPais},
		{nombre: "NumEstacion", tipo: tCatCartaPorte_cEstaciones},
		{nombre: "NombreEstacion", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NombreEstacion},
		{nombre: "NavegacionTrafico", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino_NavegacionTrafico},
		{nombre: "FechaHoraProgLlegada", tipo: tTdCFDI_tFechaH, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Domicilio
var eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Ubicacion
var eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Ubicacion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoEstacion", tipo: tCatCartaPorte_cTipoEstacion},
		{nombre: "DistanciaRecorrida", tipo: tCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_DistanciaRecorrida},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Origen, min: 0, max: 1},
		{elem: eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Destino, min: 0, max: 1},
		{elem: eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Ubicaciones
var eCartaPorte10_CartaPorte_Ubicaciones = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Ubicaciones", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Ubicaciones_Ubicacion, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}CantidadTransporta
var eCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "CantidadTransporta", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Cantidad", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta_Cantidad, requerido: true},
		{nombre: "IDOrigen", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDOrigen, requerido: true},
		{nombre: "IDDestino", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta_IDDestino, requerido: true},
		{nombre: "CvesTransporte", tipo: tCatCartaPorte_cCveTransporte},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}DetalleMercancia
var eCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "DetalleMercancia", vacio: true,
	attrs: []atributoXSD{
		{nombre: "UnidadPeso", tipo: tCatCartaPorte_cClaveUnidadPeso, requerido: true},
		{nombre: "PesoBruto", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoBruto, requerido: true},
		{nombre: "PesoNeto", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoNeto, requerido: true},
		{nombre: "PesoTara", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_PesoTara, requerido: true},
		{nombre: "NumPiezas", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia_NumPiezas},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Mercancia
var eCartaPorte10_CartaPorte_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "BienesTransp", tipo: tCatCartaPorte_cClaveProdServCP},
		{nombre: "ClaveSTCC", tipo: tCatCartaPorte_cClaveProdSTCC},
		{nombre: "Descripcion", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_Descripcion},
		{nombre: "Cantidad", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_Cantidad},
		{nombre: "ClaveUnidad", tipo: tCatCFDI_cClaveUnidad},
		{nombre: "Unidad", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_Unidad},
		{nombre: "Dimensiones", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_Dimensiones},
		{nombre: "MaterialPeligroso", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_MaterialPeligroso},
		{nombre: "CveMaterialPeligroso", tipo: tCatCartaPorte_cMaterialPeligroso},
		{nombre: "Embalaje", tipo: tCatCartaPorte_cTipoEmbalaje},
		{nombre: "DescripEmbalaje", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_DescripEmbalaje},
		{nombre: "PesoEnKg", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_PesoEnKg, requerido: true},
		{nombre: "ValorMercancia", tipo: tTdCFDI_tImporte},
		{nombre: "Moneda", tipo: tCatCFDI_cMoneda},
		{nombre: "FraccionArancelaria", tipo: tCatComExt_cFraccionArancelaria},
		{nombre: "UUIDComercioExt", tipo: tCartaPorte10_CartaPorte_Mercancias_Mercancia_UUIDComercioExt},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_Mercancia_CantidadTransporta, min: 0, max: -1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_Mercancia_DetalleMercancia, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}IdentificacionVehicular
var eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_IdentificacionVehicular = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "IdentificacionVehicular", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ConfigVehicular", tipo: tCatCartaPorte_cConfigAutotransporte, requerido: true},
		{nombre: "PlacaVM", tipo: tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_IdentificacionVehicular_PlacaVM, requerido: true},
		{nombre: "AnioModeloVM", tipo: tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_IdentificacionVehicular_AnioModeloVM, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Remolque
var eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_Remolques_Remolque = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Remolque", vacio: true,
	attrs: []atributoXSD{
		{nombre: "SubTipoRem", tipo: tCatCartaPorte_cSubTipoRem, requerido: true},
		{nombre: "Placa", tipo: tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_Remolques_Remolque_Placa, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Remolques
var eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_Remolques = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Remolques", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_Remolques_Remolque, min: 0, max: 2},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}AutotransporteFederal
var eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "AutotransporteFederal", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_NumPermisoSCT, requerido: true},
		{nombre: "NombreAseg", tipo: tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_NombreAseg, requerido: true},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_NumPolizaSeguro, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_IdentificacionVehicular, min: 1, max: 1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal_Remolques, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Contenedor
var eCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "MatriculaContenedor", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_MatriculaContenedor, requerido: true},
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedorMaritimo, requerido: true},
		{nombre: "NumPrecinto", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Contenedor_NumPrecinto},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}TransporteMaritimo
var eCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "TransporteMaritimo", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumPermisoSCT},
		{nombre: "NombreAseg", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumPolizaSeguro},
		{nombre: "TipoEmbarcacion", tipo: tCatCartaPorte_cConfigMaritima, requerido: true},
		{nombre: "Matricula", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Matricula, requerido: true},
		{nombre: "NumeroOMI", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumeroOMI, requerido: true},
		{nombre: "AnioEmbarcacion", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_AnioEmbarcacion},
		{nombre: "NombreEmbarc", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NombreEmbarc},
		{nombre: "NacionalidadEmbarc", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "UnidadesDeArqBruto", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_UnidadesDeArqBruto, requerido: true},
		{nombre: "TipoCarga", tipo: tCatCartaPorte_cClaveTipoCarga, requerido: true},
		{nombre: "NumCertITC", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumCertITC, requerido: true},
		{nombre: "Eslora", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Eslora},
		{nombre: "Manga", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Manga},
		{nombre: "Calado", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Calado},
		{nombre: "LineaNaviera", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_LineaNaviera},
		{nombre: "NombreAgenteNaviero", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NombreAgenteNaviero, requerido: true},
		{nombre: "NumAutorizacionNaviero", tipo: tCatCartaPorte_cNumAutorizacionNaviero, requerido: true},
		{nombre: "NumViaje", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumViaje},
		{nombre: "NumConocEmbarc", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_NumConocEmbarc},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo_Contenedor, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}TransporteAereo
var eCartaPorte10_CartaPorte_Mercancias_TransporteAereo = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "TransporteAereo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PermSCT", tipo: tCatCartaPorte_cTipoPermiso, requerido: true},
		{nombre: "NumPermisoSCT", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumPermisoSCT, requerido: true},
		{nombre: "MatriculaAeronave", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_MatriculaAeronave, requerido: true},
		{nombre: "NombreAseg", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumPolizaSeguro},
		{nombre: "NumeroGuia", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumeroGuia, requerido: true},
		{nombre: "LugarContrato", tipo: tTdCFDI_tDescrip120},
		{nombre: "RFCTransportista", tipo: tTdCFDI_tRFC},
		{nombre: "CodigoTransportista", tipo: tCatCartaPorte_cCodigoTransporteAereo, requerido: true},
		{nombre: "NumRegIdTribTranspor", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribTranspor},
		{nombre: "ResidenciaFiscalTranspor", tipo: tCatCFDI_cPais},
		{nombre: "NombreTransportista", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NombreTransportista},
		{nombre: "RFCEmbarcador", tipo: tTdCFDI_tRFC},
		{nombre: "NumRegIdTribEmbarc", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NumRegIdTribEmbarc},
		{nombre: "ResidenciaFiscalEmbarc", tipo: tCatCFDI_cPais},
		{nombre: "NombreEmbarcador", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteAereo_NombreEmbarcador},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}DerechosDePaso
var eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "DerechosDePaso", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDerechoDePaso", tipo: tCatCartaPorte_cDerechosDePaso, requerido: true},
		{nombre: "KilometrajePagado", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso_KilometrajePagado, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Contenedor
var eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Contenedor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoContenedor", tipo: tCatCartaPorte_cContenedor, requerido: true},
		{nombre: "PesoContenedorVacio", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoContenedorVacio, requerido: true},
		{nombre: "PesoNetoMercancia", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor_PesoNetoMercancia, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Carro
var eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Carro", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoCarro", tipo: tCatCartaPorte_cTipoCarro, requerido: true},
		{nombre: "MatriculaCarro", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_MatriculaCarro, requerido: true},
		{nombre: "GuiaCarro", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_GuiaCarro, requerido: true},
		{nombre: "ToneladasNetasCarro", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_ToneladasNetasCarro, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro_Contenedor, min: 0, max: -1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}TransporteFerroviario
var eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "TransporteFerroviario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoDeServicio", tipo: tCatCartaPorte_cTipoDeServicio, requerido: true},
		{nombre: "NombreAseg", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_NombreAseg},
		{nombre: "NumPolizaSeguro", tipo: tCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_NumPolizaSeguro},
		{nombre: "Concesionario", tipo: tTdCFDI_tRFCPM},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_DerechosDePaso, min: 0, max: -1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario_Carro, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Mercancias
var eCartaPorte10_CartaPorte_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "PesoBrutoTotal", tipo: tCartaPorte10_CartaPorte_Mercancias_PesoBrutoTotal},
		{nombre: "UnidadPeso", tipo: tCatCartaPorte_cClaveUnidadPeso},
		{nombre: "PesoNetoTotal", tipo: tCartaPorte10_CartaPorte_Mercancias_PesoNetoTotal},
		{nombre: "NumTotalMercancias", tipo: tCartaPorte10_CartaPorte_Mercancias_NumTotalMercancias, requerido: true},
		{nombre: "CargoPorTasacion", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Mercancias_Mercancia, min: 1, max: -1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_AutotransporteFederal, min: 0, max: 1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteMaritimo, min: 0, max: 1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteAereo, min: 0, max: 1},
		{elem: eCartaPorte10_CartaPorte_Mercancias_TransporteFerroviario, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Domicilio
var eCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Operador
var eCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Operador", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "RFCOperador", tipo: tTdCFDI_tRFCPF},
		{nombre: "NumLicencia", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_NumLicencia},
		{nombre: "NombreOperador", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_NombreOperador},
		{nombre: "NumRegIdTribOperador", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_NumRegIdTribOperador},
		{nombre: "ResidenciaFiscalOperador", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Operadores
var eCartaPorte10_CartaPorte_FiguraTransporte_Operadores = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Operadores", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Operadores_Operador, min: 1, max: -1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Domicilio
var eCartaPorte10_CartaPorte_FiguraTransporte_Propietario_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Propietario
var eCartaPorte10_CartaPorte_FiguraTransporte_Propietario = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Propietario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "RFCPropietario", tipo: tTdCFDI_tRFC},
		{nombre: "NombrePropietario", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_NombrePropietario},
		{nombre: "NumRegIdTribPropietario", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Propietario_NumRegIdTribPropietario},
		{nombre: "ResidenciaFiscalPropietario", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Propietario_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Domicilio
var eCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Arrendatario
var eCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Arrendatario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "RFCArrendatario", tipo: tTdCFDI_tRFC},
		{nombre: "NombreArrendatario", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_NombreArrendatario},
		{nombre: "NumRegIdTribArrendatario", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_NumRegIdTribArrendatario},
		{nombre: "ResidenciaFiscalArrendatario", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Domicilio
var eCartaPorte10_CartaPorte_FiguraTransporte_Notificado_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tTdCFDI_tDescrip100, requerido: true},
		{nombre: "NumeroExterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "NumeroInterior", tipo: tTdCFDI_tNumeroDomicilio},
		{nombre: "Colonia", tipo: tTdCFDI_tDescrip120},
		{nombre: "Localidad", tipo: tTdCFDI_tDescrip120},
		{nombre: "Referencia", tipo: tTdCFDI_tReferencia},
		{nombre: "Municipio", tipo: tTdCFDI_tDescrip120},
		{nombre: "Estado", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tCatCFDI_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}Notificado
var eCartaPorte10_CartaPorte_FiguraTransporte_Notificado = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "Notificado", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "RFCNotificado", tipo: tTdCFDI_tRFC},
		{nombre: "NombreNotificado", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_NombreNotificado},
		{nombre: "NumRegIdTribNotificado", tipo: tCartaPorte10_CartaPorte_FiguraTransporte_Notificado_NumRegIdTribNotificado},
		{nombre: "ResidenciaFiscalNotificado", tipo: tCatCFDI_cPais},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Notificado_Domicilio, min: 0, max: 1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}FiguraTransporte
var eCartaPorte10_CartaPorte_FiguraTransporte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "FiguraTransporte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "CveTransporte", tipo: tCatCartaPorte_cCveTransporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Operadores, min: 0, max: -1},
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Propietario, min: 0, max: -1},
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Arrendatario, min: 0, max: -1},
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte_Notificado, min: 0, max: -1},
	}}

// cfd/CartaPorte/CartaPorte.xsd: {http://www.sat.gob.mx/CartaPorte}CartaPorte
var eCartaPorte10_CartaPorte = &elementoXSD{ns: "http://www.sat.gob.mx/CartaPorte", nombre: "CartaPorte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TranspInternac", tipo: tCartaPorte10_CartaPorte_TranspInternac, requerido: true},
		{nombre: "EntradaSalidaMerc", tipo: tCartaPorte10_CartaPorte_EntradaSalidaMerc},
		{nombre: "ViaEntradaSalida", tipo: tCatCartaPorte_cCveTransporte},
		{nombre: "TotalDistRec", tipo: tCartaPorte10_CartaPorte_TotalDistRec},
	},
	particulas: []particulaXSD{
		{elem: eCartaPorte10_CartaPorte_Ubicaciones, min: 1, max: 1},
		{elem: eCartaPorte10_CartaPorte_Mercancias, min: 1, max: 1},
		{elem: eCartaPorte10_CartaPorte_FiguraTransporte, min: 0, max: 1},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}Emisor
var eComExt10_ComercioExterior_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "Emisor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Curp", tipo: tComExt10_tCURP},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}Receptor
var eComExt10_ComercioExterior_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "Receptor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Curp", tipo: tComExt10_tCURP},
		{nombre: "NumRegIdTrib", tipo: tComExt10_ComercioExterior_Receptor_NumRegIdTrib, requerido: true},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}Domicilio
var eComExt10_ComercioExterior_Destinatario_Domicilio = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "Domicilio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_Calle, requerido: true},
		{nombre: "NumeroExterior", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_NumeroExterior},
		{nombre: "NumeroInterior", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_NumeroInterior},
		{nombre: "Colonia", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_Colonia},
		{nombre: "Localidad", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_Localidad},
		{nombre: "Referencia", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_Referencia},
		{nombre: "Municipio", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_Municipio},
		{nombre: "Estado", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_Estado, requerido: true},
		{nombre: "Pais", tipo: tComExt10_cPais, requerido: true},
		{nombre: "CodigoPostal", tipo: tComExt10_ComercioExterior_Destinatario_Domicilio_CodigoPostal, requerido: true},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}Destinatario
var eComExt10_ComercioExterior_Destinatario = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "Destinatario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tComExt10_ComercioExterior_Destinatario_NumRegIdTrib},
		{nombre: "Rfc", tipo: tComExt10_tRFC},
		{nombre: "Curp", tipo: tComExt10_tCURP},
		{nombre: "Nombre", tipo: tComExt10_ComercioExterior_Destinatario_Nombre},
	},
	particulas: []particulaXSD{
		{elem: eComExt10_ComercioExterior_Destinatario_Domicilio, min: 1, max: 1},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}DescripcionesEspecificas
var eComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "DescripcionesEspecificas", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Marca", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Marca, requerido: true},
		{nombre: "Modelo", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_Modelo},
		{nombre: "SubModelo", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_SubModelo},
		{nombre: "NumeroSerie", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas_NumeroSerie},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}Mercancia
var eComExt10_ComercioExterior_Mercancias_Mercancia = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "Mercancia", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NoIdentificacion", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_NoIdentificacion, requerido: true},
		{nombre: "FraccionArancelaria", tipo: tComExt10_cFraccionArancelaria},
		{nombre: "CantidadAduana", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_CantidadAduana},
		{nombre: "UnidadAduana", tipo: tComExt10_cUnidadMedidaAduana},
		{nombre: "ValorUnitarioAduana", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_ValorUnitarioAduana},
		{nombre: "ValorDolares", tipo: tComExt10_ComercioExterior_Mercancias_Mercancia_ValorDolares, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eComExt10_ComercioExterior_Mercancias_Mercancia_DescripcionesEspecificas, min: 0, max: -1},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}Mercancias
var eComExt10_ComercioExterior_Mercancias = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "Mercancias", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eComExt10_ComercioExterior_Mercancias_Mercancia, min: 1, max: -1},
	}}

// cfd/ComercioExterior/ComercioExterior10.xsd: {http://www.sat.gob.mx/ComercioExterior}ComercioExterior
var eComExt10_ComercioExterior = &elementoXSD{ns: "http://www.sat.gob.mx/ComercioExterior", nombre: "ComercioExterior", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TipoOperacion", tipo: tComExt10_cTipoOperacion, requerido: true},
		{nombre: "ClaveDePedimento", tipo: tComExt10_cClavePedimento},
		{nombre: "CertificadoOrigen", tipo: tComExt10_ComercioExterior_CertificadoOrigen},
		{nombre: "NumCertificadoOrigen", tipo: tComExt10_ComercioExterior_NumCertificadoOrigen},
		{nombre: "NumeroExportadorConfiable", tipo: tComExt10_ComercioExterior_NumeroExportadorConfiable},
		{nombre: "Incoterm", tipo: tComExt10_cINCOTERM},
		{nombre: "Subdivision", tipo: tComExt10_ComercioExterior_Subdivision},
		{nombre: "Observaciones", tipo: tComExt10_ComercioExterior_Observaciones},
		{nombre: "TipoCambioUSD", tipo: tComExt10_tImporte},
		{nombre: "TotalUSD", tipo: tComExt10_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eComExt10_ComercioExterior_Emisor, min: 0, max: 1},
		{elem: eComExt10_ComercioExterior_Receptor, min: 1, max: 1},
		{elem: eComExt10_ComercioExterior_Destinatario, min: 0, max: 1},
		{elem: eComExt10_ComercioExterior_Mercancias, min: 0, max: 1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc11.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible}Traslado
var eECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Impuesto", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Impuesto, requerido: true},
		{nombre: "TasaoCuota", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_TasaoCuota, requerido: true},
		{nombre: "Importe", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_Importe, requerido: true},
	}}

// cfd/EstadoDeCuentaCombustible/ecc11.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible}Traslados
var eECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc11.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible}ConceptoEstadoDeCuentaCombustible
var eECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible", nombre: "ConceptoEstadoDeCuentaCombustible", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Identificador", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Identificador, requerido: true},
		{nombre: "Fecha", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Fecha, requerido: true},
		{nombre: "Rfc", tipo: tECC11_tRFC, requerido: true},
		{nombre: "ClaveEstacion", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ClaveEstacion, requerido: true},
		{nombre: "TAR", tipo: tECC11_cTAR},
		{nombre: "Cantidad", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Cantidad, requerido: true},
		{nombre: "NoIdentificacion", tipo: tECC11_cclaveProducto, requerido: true},
		{nombre: "Unidad", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Unidad},
		{nombre: "NombreCombustible", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_NombreCombustible, requerido: true},
		{nombre: "FolioOperacion", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_FolioOperacion, requerido: true},
		{nombre: "ValorUnitario", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_ValorUnitario, requerido: true},
		{nombre: "Importe", tipo: tECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Importe, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados, min: 1, max: 1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc11.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible}Conceptos
var eECC11_EstadoDeCuentaCombustible_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eECC11_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible, min: 1, max: -1},
	}}

// cfd/EstadoDeCuentaCombustible/ecc11.xsd: {http://www.sat.gob.mx/EstadoDeCuentaCombustible}EstadoDeCuentaCombustible
var eECC11_EstadoDeCuentaCombustible = &elementoXSD{ns: "http://www.sat.gob.mx/EstadoDeCuentaCombustible", nombre: "EstadoDeCuentaCombustible", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tECC11_EstadoDeCuentaCombustible_Version, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "TipoOperacion", tipo: tECC11_EstadoDeCuentaCombustible_TipoOperacion, requerido: true, fijo: "Tarjeta", conFijo: true},
		{nombre: "NumeroDeCuenta", tipo: tECC11_EstadoDeCuentaCombustible_NumeroDeCuenta, requerido: true},
		{nombre: "SubTotal", tipo: tECC11_EstadoDeCuentaCombustible_SubTotal, requerido: true},
		{nombre: "Total", tipo: tECC11_EstadoDeCuentaCombustible_Total, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eECC11_EstadoDeCuentaCombustible_Conceptos, min: 1, max: 1},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Percepcion
var eNomina11_Nomina_Percepciones_Percepcion = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Percepcion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoPercepcion", tipo: tNomina11_Nomina_Percepciones_Percepcion_TipoPercepcion, requerido: true},
		{nombre: "Clave", tipo: tNomina11_Nomina_Percepciones_Percepcion_Clave, requerido: true},
		{nombre: "Concepto", tipo: tNomina11_Nomina_Percepciones_Percepcion_Concepto, requerido: true},
		{nombre: "ImporteGravado", tipo: tNomina11_tImporte, requerido: true},
		{nombre: "ImporteExento", tipo: tNomina11_tImporte, requerido: true},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Percepciones
var eNomina11_Nomina_Percepciones = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Percepciones", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalGravado", tipo: tNomina11_tImporte, requerido: true},
		{nombre: "TotalExento", tipo: tNomina11_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eNomina11_Nomina_Percepciones_Percepcion, min: 1, max: -1},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Deduccion
var eNomina11_Nomina_Deducciones_Deduccion = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Deduccion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoDeduccion", tipo: tNomina11_Nomina_Deducciones_Deduccion_TipoDeduccion, requerido: true},
		{nombre: "Clave", tipo: tNomina11_Nomina_Deducciones_Deduccion_Clave, requerido: true},
		{nombre: "Concepto", tipo: tNomina11_Nomina_Deducciones_Deduccion_Concepto, requerido: true},
		{nombre: "ImporteGravado", tipo: tNomina11_tImporte, requerido: true},
		{nombre: "ImporteExento", tipo: tNomina11_tImporte, requerido: true},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Deducciones
var eNomina11_Nomina_Deducciones = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Deducciones", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TotalGravado", tipo: tNomina11_tImporte, requerido: true},
		{nombre: "TotalExento", tipo: tNomina11_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eNomina11_Nomina_Deducciones_Deduccion, min: 1, max: -1},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Incapacidad
var eNomina11_Nomina_Incapacidades_Incapacidad = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Incapacidad", vacio: true,
	attrs: []atributoXSD{
		{nombre: "DiasIncapacidad", tipo: tNomina11_Nomina_Incapacidades_Incapacidad_DiasIncapacidad, requerido: true},
		{nombre: "TipoIncapacidad", tipo: xsint, requerido: true},
		{nombre: "Descuento", tipo: tNomina11_tImporte, requerido: true},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Incapacidades
var eNomina11_Nomina_Incapacidades = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Incapacidades", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNomina11_Nomina_Incapacidades_Incapacidad, min: 1, max: -1},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}HorasExtra
var eNomina11_Nomina_HorasExtras_HorasExtra = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "HorasExtra", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Dias", tipo: xsint, requerido: true},
		{nombre: "TipoHoras", tipo: tNomina11_Nomina_HorasExtras_HorasExtra_TipoHoras, requerido: true},
		{nombre: "HorasExtra", tipo: xsint, requerido: true},
		{nombre: "ImportePagado", tipo: tNomina11_tImporte, requerido: true},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}HorasExtras
var eNomina11_Nomina_HorasExtras = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "HorasExtras", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNomina11_Nomina_HorasExtras_HorasExtra, min: 1, max: -1},
	}}

// cfd/nomina/nomina11.xsd: {http://www.sat.gob.mx/nomina}Nomina
var eNomina11_Nomina = &elementoXSD{ns: "http://www.sat.gob.mx/nomina", nombre: "Nomina", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "RegistroPatronal", tipo: tNomina11_Nomina_RegistroPatronal},
		{nombre: "NumEmpleado", tipo: tNomina11_Nomina_NumEmpleado, requerido: true},
		{nombre: "CURP", tipo: tNomina11_tCURP, requerido: true},
		{nombre: "TipoRegimen", tipo: tNomina11_Nomina_TipoRegimen, requerido: true},
		{nombre: "NumSeguridadSocial", tipo: tNomina11_Nomina_NumSeguridadSocial},
		{nombre: "FechaPago", tipo: xsdate, requerido: true},
		{nombre: "FechaInicialPago", tipo: xsdate, requerido: true},
		{nombre: "FechaFinalPago", tipo: xsdate, requerido: true},
		{nombre: "NumDiasPagados", tipo: tNomina11_Nomina_NumDiasPagados, requerido: true},
		{nombre: "Departamento", tipo: tNomina11_Nomina_Departamento},
		{nombre: "CLABE", tipo: tNomina11_tClabe},
		{nombre: "Banco", tipo: tNomina11_Nomina_Banco},
		{nombre: "FechaInicioRelLaboral", tipo: xsdate},
		{nombre: "Antiguedad", tipo: xsint},
		{nombre: "Puesto", tipo: xsstring},
		{nombre: "TipoContrato", tipo: xsstring},
		{nombre: "TipoJornada", tipo: xsstring},
		{nombre: "PeriodicidadPago", tipo: tNomina11_Nomina_PeriodicidadPago, requerido: true},
		{nombre: "SalarioBaseCotApor", tipo: tNomina11_tImporte},
		{nombre: "RiesgoPuesto", tipo: tNomina11_Nomina_RiesgoPuesto},
		{nombre: "SalarioDiarioIntegrado", tipo: tNomina11_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eNomina11_Nomina_Percepciones, min: 0, max: 1},
		{elem: eNomina11_Nomina_Deducciones, min: 0, max: 1},
		{elem: eNomina11_Nomina_Incapacidades, min: 0, max: 1},
		{elem: eNomina11_Nomina_HorasExtras, min: 0, max: 1},
	}}

// cfd/ine/ine11.xsd: {http://www.sat.gob.mx/ine}Contabilidad
var eINE11_INE_Entidad_Contabilidad = &elementoXSD{ns: "http://www.sat.gob.mx/ine", nombre: "Contabilidad", vacio: true,
	attrs: []atributoXSD{
		{nombre: "IdContabilidad", tipo: tINE11_INE_Entidad_Contabilidad_IdContabilidad, requerido: true},
	}}

// cfd/ine/ine11.xsd: {http://www.sat.gob.mx/ine}Entidad
var eINE11_INE_Entidad = &elementoXSD{ns: "http://www.sat.gob.mx/ine", nombre: "Entidad", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ClaveEntidad", tipo: tINE11_tClaveEntidad, requerido: true},
		{nombre: "Ambito", tipo: tINE11_INE_Entidad_Ambito},
	},
	particulas: []particulaXSD{
		{elem: eINE11_INE_Entidad_Contabilidad, min: 0, max: -1},
	}}

// cfd/ine/ine11.xsd: {http://www.sat.gob.mx/ine}INE
var eINE11_INE = &elementoXSD{ns: "http://www.sat.gob.mx/ine", nombre: "INE", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tINE11_INE_Version, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "TipoProceso", tipo: tINE11_INE_TipoProceso, requerido: true},
		{nombre: "TipoComite", tipo: tINE11_INE_TipoComite},
		{nombre: "IdContabilidad", tipo: tINE11_INE_IdContabilidad},
	},
	particulas: []particulaXSD{
		{elem: eINE11_INE_Entidad, min: 0, max: 76},
	}}

// cfd/valesdedespensa/valesdedespensa.xsd: {http://www.sat.gob.mx/valesdedespensa}Concepto
var eValesDespensa_ValesDeDespensa_Conceptos_Concepto = &elementoXSD{ns: "http://www.sat.gob.mx/valesdedespensa", nombre: "Concepto", vacio: true,
	attrs: []atributoXSD{
		{nombre: "identificador", tipo: tValesDespensa_ValesDeDespensa_Conceptos_Concepto_identificador, requerido: true},
		{nombre: "fecha", tipo: tValesDespensa_ValesDeDespensa_Conceptos_Concepto_fecha, requerido: true},
		{nombre: "rfc", tipo: tValesDespensa_tRFC, requerido: true},
		{nombre: "curp", tipo: tValesDespensa_tCURP, requerido: true},
		{nombre: "nombre", tipo: tValesDespensa_ValesDeDespensa_Conceptos_Concepto_nombre, requerido: true},
		{nombre: "numSeguridadSocial", tipo: tValesDespensa_ValesDeDespensa_Conceptos_Concepto_numSeguridadSocial},
		{nombre: "importe", tipo: tValesDespensa_ValesDeDespensa_Conceptos_Concepto_importe, requerido: true},
	}}

// cfd/valesdedespensa/valesdedespensa.xsd: {http://www.sat.gob.mx/valesdedespensa}Conceptos
var eValesDespensa_ValesDeDespensa_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/valesdedespensa", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eValesDespensa_ValesDeDespensa_Conceptos_Concepto, min: 1, max: -1},
	}}

// cfd/valesdedespensa/valesdedespensa.xsd: {http://www.sat.gob.mx/valesdedespensa}ValesDeDespensa
var eValesDespensa_ValesDeDespensa = &elementoXSD{ns: "http://www.sat.gob.mx/valesdedespensa", nombre: "ValesDeDespensa", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "tipoOperacion", tipo: tValesDespensa_ValesDeDespensa_tipoOperacion, requerido: true, fijo: "monedero electrónico", conFijo: true},
		{nombre: "registroPatronal", tipo: tValesDespensa_ValesDeDespensa_registroPatronal},
		{nombre: "numeroDeCuenta", tipo: tValesDespensa_ValesDeDespensa_numeroDeCuenta, requerido: true},
		{nombre: "total", tipo: tValesDespensa_ValesDeDespensa_total, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eValesDespensa_ValesDeDespensa_Conceptos, min: 1, max: 1},
	}}

// cfd/consumodecombustibles/consumodecombustibles.xsd: {http://www.sat.gob.mx/consumodecombustibles}Determinado
var eConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado = &elementoXSD{ns: "http://www.sat.gob.mx/consumodecombustibles", nombre: "Determinado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "impuesto", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_impuesto, requerido: true},
		{nombre: "tasa", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_tasa, requerido: true},
		{nombre: "importe", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_importe, requerido: true},
	}}

// cfd/consumodecombustibles/consumodecombustibles.xsd: {http://www.sat.gob.mx/consumodecombustibles}Determinados
var eConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados = &elementoXSD{ns: "http://www.sat.gob.mx/consumodecombustibles", nombre: "Determinados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado, min: 1, max: -1},
	}}

// cfd/consumodecombustibles/consumodecombustibles.xsd: {http://www.sat.gob.mx/consumodecombustibles}ConceptoConsumoDeCombustibles
var eConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles = &elementoXSD{ns: "http://www.sat.gob.mx/consumodecombustibles", nombre: "ConceptoConsumoDeCombustibles", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "identificador", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_identificador, requerido: true},
		{nombre: "fecha", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_fecha, requerido: true},
		{nombre: "rfc", tipo: tConsumoComb10_tRFC, requerido: true},
		{nombre: "claveEstacion", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_claveEstacion, requerido: true},
		{nombre: "cantidad", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_cantidad, requerido: true},
		{nombre: "nombreCombustible", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_nombreCombustible, requerido: true},
		{nombre: "folioOperacion", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_folioOperacion, requerido: true},
		{nombre: "valorUnitario", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_valorUnitario, requerido: true},
		{nombre: "importe", tipo: tConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_importe, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados, min: 1, max: 1},
	}}

// cfd/consumodecombustibles/consumodecombustibles.xsd: {http://www.sat.gob.mx/consumodecombustibles}Conceptos
var eConsumoComb10_ConsumoDeCombustibles_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/consumodecombustibles", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eConsumoComb10_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles, min: 1, max: -1},
	}}

// cfd/consumodecombustibles/consumodecombustibles.xsd: {http://www.sat.gob.mx/consumodecombustibles}ConsumoDeCombustibles
var eConsumoComb10_ConsumoDeCombustibles = &elementoXSD{ns: "http://www.sat.gob.mx/consumodecombustibles", nombre: "ConsumoDeCombustibles", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "tipoOperacion", tipo: tConsumoComb10_ConsumoDeCombustibles_tipoOperacion, requerido: true, fijo: "monedero electrónico", conFijo: true},
		{nombre: "numeroDeCuenta", tipo: tConsumoComb10_ConsumoDeCombustibles_numeroDeCuenta, requerido: true},
		{nombre: "subTotal", tipo: tConsumoComb10_ConsumoDeCombustibles_subTotal},
		{nombre: "total", tipo: tConsumoComb10_ConsumoDeCombustibles_total, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eConsumoComb10_ConsumoDeCombustibles_Conceptos, min: 1, max: 1},
	}}

// cfd/consumodecombustibles/consumodeCombustibles11.xsd: {http://www.sat.gob.mx/ConsumoDeCombustibles11}Determinado
var eConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado = &elementoXSD{ns: "http://www.sat.gob.mx/ConsumoDeCombustibles11", nombre: "Determinado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "impuesto", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_impuesto, requerido: true},
		{nombre: "tasaOCuota", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_tasaOCuota, requerido: true},
		{nombre: "importe", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado_importe, requerido: true},
	}}

// cfd/consumodecombustibles/consumodeCombustibles11.xsd: {http://www.sat.gob.mx/ConsumoDeCombustibles11}Determinados
var eConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados = &elementoXSD{ns: "http://www.sat.gob.mx/ConsumoDeCombustibles11", nombre: "Determinados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados_Determinado, min: 1, max: -1},
	}}

// cfd/consumodecombustibles/consumodeCombustibles11.xsd: {http://www.sat.gob.mx/ConsumoDeCombustibles11}ConceptoConsumoDeCombustibles
var eConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles = &elementoXSD{ns: "http://www.sat.gob.mx/ConsumoDeCombustibles11", nombre: "ConceptoConsumoDeCombustibles", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "identificador", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_identificador, requerido: true},
		{nombre: "fecha", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "rfc", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "claveEstacion", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_claveEstacion, requerido: true},
		{nombre: "tipoCombustible", tipo: tCatCombustible_cClaveTipoCombustible, requerido: true},
		{nombre: "cantidad", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_cantidad, requerido: true},
		{nombre: "nombreCombustible", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_nombreCombustible, requerido: true},
		{nombre: "folioOperacion", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_folioOperacion, requerido: true},
		{nombre: "valorUnitario", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_valorUnitario, requerido: true},
		{nombre: "importe", tipo: tConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_importe, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles_Determinados, min: 1, max: 1},
	}}

// cfd/consumodecombustibles/consumodeCombustibles11.xsd: {http://www.sat.gob.mx/ConsumoDeCombustibles11}Conceptos
var eConsumoComb11_ConsumoDeCombustibles_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/ConsumoDeCombustibles11", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eConsumoComb11_ConsumoDeCombustibles_Conceptos_ConceptoConsumoDeCombustibles, min: 1, max: -1},
	}}

// cfd/consumodecombustibles/consumodeCombustibles11.xsd: {http://www.sat.gob.mx/ConsumoDeCombustibles11}ConsumoDeCombustibles
var eConsumoComb11_ConsumoDeCombustibles = &elementoXSD{ns: "http://www.sat.gob.mx/ConsumoDeCombustibles11", nombre: "ConsumoDeCombustibles", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: tConsumoComb11_ConsumoDeCombustibles_version, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "tipoOperacion", tipo: tConsumoComb11_ConsumoDeCombustibles_tipoOperacion, requerido: true, fijo: "monedero electrónico", conFijo: true},
		{nombre: "numeroDeCuenta", tipo: tConsumoComb11_ConsumoDeCombustibles_numeroDeCuenta, requerido: true},
		{nombre: "subTotal", tipo: tConsumoComb11_ConsumoDeCombustibles_subTotal},
		{nombre: "total", tipo: tConsumoComb11_ConsumoDeCombustibles_total, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eConsumoComb11_ConsumoDeCombustibles_Conceptos, min: 1, max: 1},
	}}

// cfd/pagoenespecie/pagoenespecie.xsd: {http://www.sat.gob.mx/pagoenespecie}PagoEnEspecie
var ePagoEnEspecie_PagoEnEspecie = &elementoXSD{ns: "http://www.sat.gob.mx/pagoenespecie", nombre: "PagoEnEspecie", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "CvePIC", tipo: tPagoEnEspecie_PagoEnEspecie_CvePIC, requerido: true},
		{nombre: "FolioSolDon", tipo: tPagoEnEspecie_PagoEnEspecie_FolioSolDon, requerido: true},
		{nombre: "PzaArtNombre", tipo: tPagoEnEspecie_PagoEnEspecie_PzaArtNombre, requerido: true},
		{nombre: "PzaArtTecn", tipo: tPagoEnEspecie_PagoEnEspecie_PzaArtTecn, requerido: true},
		{nombre: "PzaArtAProd", tipo: tPagoEnEspecie_PagoEnEspecie_PzaArtAProd, requerido: true},
		{nombre: "PzaArtDim", tipo: tPagoEnEspecie_PagoEnEspecie_PzaArtDim, requerido: true},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DescInmueble
var eNotarios_NotariosPublicos_DescInmuebles_DescInmueble = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DescInmueble", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoInmueble", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_TipoInmueble, requerido: true},
		{nombre: "Calle", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Calle, requerido: true},
		{nombre: "NoExterior", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_NoExterior},
		{nombre: "NoInterior", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_NoInterior},
		{nombre: "Colonia", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Colonia},
		{nombre: "Localidad", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Localidad},
		{nombre: "Referencia", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Referencia},
		{nombre: "Municipio", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Municipio, requerido: true},
		{nombre: "Estado", tipo: tNotarios_tEntidadFederativa, requerido: true},
		{nombre: "Pais", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_Pais, requerido: true},
		{nombre: "CodigoPostal", tipo: tNotarios_NotariosPublicos_DescInmuebles_DescInmueble_CodigoPostal, requerido: true},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DescInmuebles
var eNotarios_NotariosPublicos_DescInmuebles = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DescInmuebles", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNotarios_NotariosPublicos_DescInmuebles_DescInmueble, min: 1, max: -1},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosOperacion
var eNotarios_NotariosPublicos_DatosOperacion = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosOperacion", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumInstrumentoNotarial", tipo: tNotarios_NotariosPublicos_DatosOperacion_NumInstrumentoNotarial, requerido: true},
		{nombre: "FechaInstNotarial", tipo: xsdate, requerido: true},
		{nombre: "MontoOperacion", tipo: tNotarios_tImporte, requerido: true},
		{nombre: "Subtotal", tipo: tNotarios_tImporte, requerido: true},
		{nombre: "IVA", tipo: tNotarios_tImporte, requerido: true},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosNotario
var eNotarios_NotariosPublicos_DatosNotario = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosNotario", vacio: true,
	attrs: []atributoXSD{
		{nombre: "CURP", tipo: tNotarios_tCURP, requerido: true},
		{nombre: "NumNotaria", tipo: tNotarios_NotariosPublicos_DatosNotario_NumNotaria, requerido: true},
		{nombre: "EntidadFederativa", tipo: tNotarios_tEntidadFederativa, requerido: true},
		{nombre: "Adscripcion", tipo: tNotarios_NotariosPublicos_DatosNotario_Adscripcion},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosUnEnajenante
var eNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosUnEnajenante", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Nombre", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante_Nombre, requerido: true},
		{nombre: "ApellidoPaterno", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante_ApellidoPaterno, requerido: true},
		{nombre: "ApellidoMaterno", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante_ApellidoMaterno},
		{nombre: "RFC", tipo: tNotarios_tRFC, requerido: true},
		{nombre: "CURP", tipo: tNotarios_tCURP, requerido: true},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosEnajenanteCopSC
var eNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosEnajenanteCopSC", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Nombre", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_Nombre, requerido: true},
		{nombre: "ApellidoPaterno", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_ApellidoPaterno},
		{nombre: "ApellidoMaterno", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_ApellidoMaterno},
		{nombre: "RFC", tipo: tNotarios_tRFC, requerido: true},
		{nombre: "CURP", tipo: tNotarios_tCURP},
		{nombre: "Porcentaje", tipo: tNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC_Porcentaje, requerido: true},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosEnajenantesCopSC
var eNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosEnajenantesCopSC", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC_DatosEnajenanteCopSC, min: 1, max: -1},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosEnajenante
var eNotarios_NotariosPublicos_DatosEnajenante = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosEnajenante", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "CoproSocConyugalE", tipo: tNotarios_NotariosPublicos_DatosEnajenante_CoproSocConyugalE, requerido: true},
	},
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eNotarios_NotariosPublicos_DatosEnajenante_DatosUnEnajenante, min: 0, max: 1},
			{elem: eNotarios_NotariosPublicos_DatosEnajenante_DatosEnajenantesCopSC, min: 0, max: 1},
		}},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosUnAdquiriente
var eNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosUnAdquiriente", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Nombre", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente_Nombre, requerido: true},
		{nombre: "ApellidoPaterno", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente_ApellidoPaterno},
		{nombre: "ApellidoMaterno", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente_ApellidoMaterno},
		{nombre: "RFC", tipo: tNotarios_tRFC, requerido: true},
		{nombre: "CURP", tipo: tNotarios_tCURP},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosAdquirienteCopSC
var eNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosAdquirienteCopSC", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Nombre", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_Nombre, requerido: true},
		{nombre: "ApellidoPaterno", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_ApellidoPaterno},
		{nombre: "ApellidoMaterno", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_ApellidoMaterno},
		{nombre: "RFC", tipo: tNotarios_tRFC, requerido: true},
		{nombre: "CURP", tipo: tNotarios_tCURP},
		{nombre: "Porcentaje", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC_Porcentaje, requerido: true},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosAdquirientesCopSC
var eNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosAdquirientesCopSC", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC_DatosAdquirienteCopSC, min: 1, max: -1},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}DatosAdquiriente
var eNotarios_NotariosPublicos_DatosAdquiriente = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "DatosAdquiriente", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "CoproSocConyugalE", tipo: tNotarios_NotariosPublicos_DatosAdquiriente_CoproSocConyugalE, requerido: true},
	},
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eNotarios_NotariosPublicos_DatosAdquiriente_DatosUnAdquiriente, min: 0, max: 1},
			{elem: eNotarios_NotariosPublicos_DatosAdquiriente_DatosAdquirientesCopSC, min: 0, max: 1},
		}},
	}}

// cfd/notariospublicos/notariospublicos.xsd: {http://www.sat.gob.mx/notariospublicos}NotariosPublicos
var eNotarios_NotariosPublicos = &elementoXSD{ns: "http://www.sat.gob.mx/notariospublicos", nombre: "NotariosPublicos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
	},
	particulas: []particulaXSD{
		{elem: eNotarios_NotariosPublicos_DescInmuebles, min: 1, max: 1},
		{elem: eNotarios_NotariosPublicos_DatosOperacion, min: 1, max: 1},
		{elem: eNotarios_NotariosPublicos_DatosNotario, min: 1, max: 1},
		{elem: eNotarios_NotariosPublicos_DatosEnajenante, min: 1, max: 1},
		{elem: eNotarios_NotariosPublicos_DatosAdquiriente, min: 1, max: 1},
	}}

// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xsd: {http://www.sat.gob.mx/servicioparcialconstruccion}Inmueble
var eParcialesConstruccion_parcialesconstruccion_Inmueble = &elementoXSD{ns: "http://www.sat.gob.mx/servicioparcialconstruccion", nombre: "Inmueble", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Calle", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_Calle, requerido: true},
		{nombre: "NoExterior", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_NoExterior},
		{nombre: "NoInterior", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_NoInterior},
		{nombre: "Colonia", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_Colonia},
		{nombre: "Localidad", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_Localidad},
		{nombre: "Referencia", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_Referencia},
		{nombre: "Municipio", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_Municipio, requerido: true},
		{nombre: "Estado", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_Estado, requerido: true},
		{nombre: "CodigoPostal", tipo: tParcialesConstruccion_parcialesconstruccion_Inmueble_CodigoPostal, requerido: true},
	}}

// cfd/servicioparcialconstruccion/servicioparcialconstruccion.xsd: {http://www.sat.gob.mx/servicioparcialconstruccion}parcialesconstruccion
var eParcialesConstruccion_parcialesconstruccion = &elementoXSD{ns: "http://www.sat.gob.mx/servicioparcialconstruccion", nombre: "parcialesconstruccion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tParcialesConstruccion_parcialesconstruccion_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "NumPerLicoAut", tipo: tParcialesConstruccion_parcialesconstruccion_NumPerLicoAut, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eParcialesConstruccion_parcialesconstruccion_Inmueble, min: 1, max: 1},
	}}

// cfd/vehiculousado/vehiculousado.xsd: {http://www.sat.gob.mx/vehiculousado}InformacionAduanera
var eVehiculoUsado_VehiculoUsado_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/vehiculousado", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tVehiculoUsado_VehiculoUsado_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tVehiculoUsado_VehiculoUsado_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tVehiculoUsado_VehiculoUsado_InformacionAduanera_aduana},
	}}

// cfd/vehiculousado/vehiculousado.xsd: {http://www.sat.gob.mx/vehiculousado}VehiculoUsado
var eVehiculoUsado_VehiculoUsado = &elementoXSD{ns: "http://www.sat.gob.mx/vehiculousado", nombre: "VehiculoUsado", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tVehiculoUsado_VehiculoUsado_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "montoAdquisicion", tipo: tVehiculoUsado_tImporte, requerido: true},
		{nombre: "montoEnajenacion", tipo: tVehiculoUsado_tImporte, requerido: true},
		{nombre: "claveVehicular", tipo: tVehiculoUsado_VehiculoUsado_claveVehicular, requerido: true},
		{nombre: "marca", tipo: tVehiculoUsado_VehiculoUsado_marca, requerido: true},
		{nombre: "tipo", tipo: tVehiculoUsado_VehiculoUsado_tipo, requerido: true},
		{nombre: "modelo", tipo: tVehiculoUsado_VehiculoUsado_modelo, requerido: true},
		{nombre: "numeroMotor", tipo: tVehiculoUsado_VehiculoUsado_numeroMotor},
		{nombre: "numeroSerie", tipo: tVehiculoUsado_VehiculoUsado_numeroSerie},
		{nombre: "NIV", tipo: tVehiculoUsado_VehiculoUsado_NIV},
		{nombre: "valor", tipo: tVehiculoUsado_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eVehiculoUsado_VehiculoUsado_InformacionAduanera, min: 0, max: -1},
	}}

// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xsd: {http://www.sat.gob.mx/TuristaPasajeroExtranjero}datosTransito
var eTuristaPasajero_TuristaPasajeroExtranjero_datosTransito = &elementoXSD{ns: "http://www.sat.gob.mx/TuristaPasajeroExtranjero", nombre: "datosTransito", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Via", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_Via, requerido: true},
		{nombre: "TipoId", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_TipoId, requerido: true},
		{nombre: "NumeroId", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_NumeroId, requerido: true},
		{nombre: "Nacionalidad", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_Nacionalidad, requerido: true},
		{nombre: "EmpresaTransporte", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_datosTransito_EmpresaTransporte, requerido: true},
		{nombre: "IdTransporte", tipo: xsstring},
	}}

// cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xsd: {http://www.sat.gob.mx/TuristaPasajeroExtranjero}TuristaPasajeroExtranjero
var eTuristaPasajero_TuristaPasajeroExtranjero = &elementoXSD{ns: "http://www.sat.gob.mx/TuristaPasajeroExtranjero", nombre: "TuristaPasajeroExtranjero", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "fechadeTransito", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_fechadeTransito, requerido: true},
		{nombre: "tipoTransito", tipo: tTuristaPasajero_TuristaPasajeroExtranjero_tipoTransito, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eTuristaPasajero_TuristaPasajeroExtranjero_datosTransito, min: 1, max: 1},
	}}

// cfd/acreditamiento/AcreditamientoIEPS10.xsd: {http://www.sat.gob.mx/acreditamiento}acreditamientoIEPS
var eAcreditamientoIEPS_acreditamientoIEPS = &elementoXSD{ns: "http://www.sat.gob.mx/acreditamiento", nombre: "acreditamientoIEPS", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TAR", tipo: tAcreditamientoIEPS_acreditamientoIEPS_TAR, requerido: true},
	}}

// cfd/spei/spei.xsd: {http://www.sat.gob.mx/spei}Ordenante
var eSPEI_ComplementoSPEI_SPEITercero_Ordenante = &elementoXSD{ns: "http://www.sat.gob.mx/spei", nombre: "Ordenante", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BancoEmisor", tipo: tSPEI_ComplementoSPEI_SPEITercero_Ordenante_BancoEmisor, requerido: true},
		{nombre: "Nombre", tipo: tSPEI_ComplementoSPEI_SPEITercero_Ordenante_Nombre, requerido: true},
		{nombre: "TipoCuenta", tipo: tSPEI_ComplementoSPEI_SPEITercero_Ordenante_TipoCuenta, requerido: true},
		{nombre: "Cuenta", tipo: tSPEI_ComplementoSPEI_SPEITercero_Ordenante_Cuenta, requerido: true},
		{nombre: "RFC", tipo: tSPEI_tRFC, requerido: true},
	}}

// cfd/spei/spei.xsd: {http://www.sat.gob.mx/spei}Beneficiario
var eSPEI_ComplementoSPEI_SPEITercero_Beneficiario = &elementoXSD{ns: "http://www.sat.gob.mx/spei", nombre: "Beneficiario", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BancoReceptor", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_BancoReceptor, requerido: true},
		{nombre: "Nombre", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_Nombre, requerido: true},
		{nombre: "TipoCuenta", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_TipoCuenta, requerido: true},
		{nombre: "Cuenta", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_Cuenta, requerido: true},
		{nombre: "RFC", tipo: tSPEI_tRFC, requerido: true},
		{nombre: "Concepto", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_Concepto, requerido: true},
		{nombre: "IVA", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_IVA},
		{nombre: "MontoPago", tipo: tSPEI_ComplementoSPEI_SPEITercero_Beneficiario_MontoPago, requerido: true},
	}}

// cfd/spei/spei.xsd: {http://www.sat.gob.mx/spei}SPEI_Tercero
var eSPEI_ComplementoSPEI_SPEITercero = &elementoXSD{ns: "http://www.sat.gob.mx/spei", nombre: "SPEI_Tercero", seqMin: 1, seqMax: 1, todos: true,
	attrs: []atributoXSD{
		{nombre: "FechaOperacion", tipo: tSPEI_ComplementoSPEI_SPEITercero_FechaOperacion, requerido: true},
		{nombre: "Hora", tipo: tSPEI_ComplementoSPEI_SPEITercero_Hora, requerido: true},
		{nombre: "ClaveSPEI", tipo: tSPEI_ComplementoSPEI_SPEITercero_ClaveSPEI, requerido: true},
		{nombre: "sello", tipo: tSPEI_ComplementoSPEI_SPEITercero_sello, requerido: true},
		{nombre: "numeroCertificado", tipo: tSPEI_ComplementoSPEI_SPEITercero_numeroCertificado, requerido: true},
		{nombre: "cadenaCDA", tipo: tSPEI_ComplementoSPEI_SPEITercero_cadenaCDA, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eSPEI_ComplementoSPEI_SPEITercero_Ordenante, min: 1, max: 1},
		{elem: eSPEI_ComplementoSPEI_SPEITercero_Beneficiario, min: 1, max: 1},
	}}

// cfd/spei/spei.xsd: {http://www.sat.gob.mx/spei}Complemento_SPEI
var eSPEI_ComplementoSPEI = &elementoXSD{ns: "http://www.sat.gob.mx/spei", nombre: "Complemento_SPEI", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eSPEI_ComplementoSPEI_SPEITercero, min: 1, max: -1},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}VehiculosUsadosEnajenadoPermAlFab
var eRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "VehiculosUsadosEnajenadoPermAlFab", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PrecioVehUsado", tipo: tRenovacion_tImporte, requerido: true},
		{nombre: "TipoVeh", tipo: tRenovacion_cTipoVehiculor, requerido: true},
		{nombre: "Marca", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Marca, requerido: true},
		{nombre: "TipooClase", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_TipooClase, requerido: true},
		{nombre: "Año", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Anio, requerido: true},
		{nombre: "Modelo", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Modelo},
		{nombre: "NIV", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NIV},
		{nombre: "NumSerie", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumSerie},
		{nombre: "NumPlacas", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumPlacas, requerido: true},
		{nombre: "NumMotor", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumMotor},
		{nombre: "NumFolTarjCir", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumFolTarjCir, requerido: true},
		{nombre: "NumPedIm", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_NumPedIm},
		{nombre: "Aduana", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Aduana},
		{nombre: "FechaRegulVeh", tipo: xsdate},
		{nombre: "Foliofiscal", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab_Foliofiscal, requerido: true},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}VehiculoNuvoSemEnajenadoFabAlPerm
var eRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "VehiculoNuvoSemEnajenadoFabAlPerm", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Año", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Anio, requerido: true},
		{nombre: "Modelo", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Modelo},
		{nombre: "NumPlacas", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_NumPlacas, requerido: true},
		{nombre: "RFC", tipo: tRenovacion_tRFC},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}DecretoRenovVehicular
var eRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "DecretoRenovVehicular", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "VehEnaj", tipo: tRenovacion_cVehiculoEnajenado, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculosUsadosEnajenadoPermAlFab, min: 1, max: -1},
		{elem: eRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular_VehiculoNuvoSemEnajenadoFabAlPerm, min: 1, max: 1},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}VehiculoUsadoEnajenadoPermAlFab
var eRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "VehiculoUsadoEnajenadoPermAlFab", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PrecioVehUsado", tipo: tRenovacion_tImporte, requerido: true},
		{nombre: "TipoVeh", tipo: tRenovacion_cTipoVehiculos, requerido: true},
		{nombre: "Marca", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Marca, requerido: true},
		{nombre: "TipooClase", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_TipooClase, requerido: true},
		{nombre: "Año", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Anio, requerido: true},
		{nombre: "Modelo", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Modelo},
		{nombre: "NIV", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NIV},
		{nombre: "NumSerie", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumSerie},
		{nombre: "NumPlacas", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumPlacas, requerido: true},
		{nombre: "NumMotor", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumMotor},
		{nombre: "NumFolTarjCir", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumFolTarjCir, requerido: true},
		{nombre: "NumFolAvisoint", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumFolAvisoint, requerido: true},
		{nombre: "NumPedIm", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_NumPedIm, requerido: true},
		{nombre: "Aduana", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Aduana, requerido: true},
		{nombre: "FechaRegulVeh", tipo: xsdate, requerido: true},
		{nombre: "Foliofiscal", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab_Foliofiscal, requerido: true},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}VehiculoNuvoSemEnajenadoFabAlPerm
var eRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "VehiculoNuvoSemEnajenadoFabAlPerm", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Año", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Anio, requerido: true},
		{nombre: "Modelo", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_Modelo},
		{nombre: "NumPlacas", tipo: tRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm_NumPlacas, requerido: true},
		{nombre: "RFC", tipo: tRenovacion_tRFC},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}DecretoSustitVehicular
var eRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "DecretoSustitVehicular", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "VehEnaj", tipo: tRenovacion_cVehiculoEnajenado, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoUsadoEnajenadoPermAlFab, min: 1, max: 1},
		{elem: eRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular_VehiculoNuvoSemEnajenadoFabAlPerm, min: 1, max: 1},
	}}

// cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd: {http://www.sat.gob.mx/renovacionysustitucionvehiculos}renovacionysustitucionvehiculos
var eRenovacion_renovacionysustitucionvehiculos = &elementoXSD{ns: "http://www.sat.gob.mx/renovacionysustitucionvehiculos", nombre: "renovacionysustitucionvehiculos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tRenovacion_renovacionysustitucionvehiculos_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TipoDeDecreto", tipo: tRenovacion_cTipoDecreto, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRenovacion_renovacionysustitucionvehiculos_DecretoRenovVehicular, min: 0, max: 1},
		{elem: eRenovacion_renovacionysustitucionvehiculos_DecretoSustitVehicular, min: 0, max: 1},
	}}

// cfd/certificadodestruccion/certificadodedestruccion.xsd: {http://www.sat.gob.mx/certificadodestruccion}VehiculoDestruido
var eCertDestruccion_certificadodedestruccion_VehiculoDestruido = &elementoXSD{ns: "http://www.sat.gob.mx/certificadodestruccion", nombre: "VehiculoDestruido", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Marca", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_Marca, requerido: true},
		{nombre: "TipooClase", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_TipooClase, requerido: true},
		{nombre: "Año", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_Anio, requerido: true},
		{nombre: "Modelo", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_Modelo},
		{nombre: "NIV", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NIV},
		{nombre: "NumSerie", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumSerie},
		{nombre: "NumPlacas", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumPlacas, requerido: true},
		{nombre: "NumMotor", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumMotor},
		{nombre: "NumFolTarjCir", tipo: tCertDestruccion_certificadodedestruccion_VehiculoDestruido_NumFolTarjCir, requerido: true},
	}}

// cfd/certificadodestruccion/certificadodedestruccion.xsd: {http://www.sat.gob.mx/certificadodestruccion}InformacionAduanera
var eCertDestruccion_certificadodedestruccion_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/certificadodestruccion", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumPedImp", tipo: tCertDestruccion_certificadodedestruccion_InformacionAduanera_NumPedImp, requerido: true},
		{nombre: "Fecha", tipo: xsdate, requerido: true},
		{nombre: "Aduana", tipo: tCertDestruccion_certificadodedestruccion_InformacionAduanera_Aduana, requerido: true},
	}}

// cfd/certificadodestruccion/certificadodedestruccion.xsd: {http://www.sat.gob.mx/certificadodestruccion}certificadodedestruccion
var eCertDestruccion_certificadodedestruccion = &elementoXSD{ns: "http://www.sat.gob.mx/certificadodestruccion", nombre: "certificadodedestruccion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tCertDestruccion_certificadodedestruccion_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "Serie", tipo: tCertDestruccion_cTipoSerie, requerido: true},
		{nombre: "NumFolDesVeh", tipo: tCertDestruccion_certificadodedestruccion_NumFolDesVeh, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eCertDestruccion_certificadodedestruccion_VehiculoDestruido, min: 1, max: 1},
		{elem: eCertDestruccion_certificadodedestruccion_InformacionAduanera, min: 0, max: 1},
	}}

// cfd/arteantiguedades/obrasarteantiguedades.xsd: {http://www.sat.gob.mx/arteantiguedades}obrasarteantiguedades
var eObrasArte_obrasarteantiguedades = &elementoXSD{ns: "http://www.sat.gob.mx/arteantiguedades", nombre: "obrasarteantiguedades", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TipoBien", tipo: tObrasArte_obrasarteantiguedades_TipoBien, requerido: true},
		{nombre: "OtrosTipoBien", tipo: tObrasArte_obrasarteantiguedades_OtrosTipoBien},
		{nombre: "TituloAdquirido", tipo: tObrasArte_obrasarteantiguedades_TituloAdquirido, requerido: true},
		{nombre: "OtrosTituloAdquirido", tipo: tObrasArte_obrasarteantiguedades_OtrosTituloAdquirido},
		{nombre: "Subtotal", tipo: tObrasArte_tImporte},
		{nombre: "IVA", tipo: tObrasArte_tImporte},
		{nombre: "FechaAdquisicion", tipo: tObrasArte_obrasarteantiguedades_FechaAdquisicion, requerido: true},
		{nombre: "CaracterísticasDeObraoPieza", tipo: tObrasArte_obrasarteantiguedades_CaracteristicasDeObraoPieza, requerido: true},
	}}

// cfd/pfic/pfic.xsd: {http://www.sat.gob.mx/pfic}PFintegranteCoordinado
var ePFIC_PFintegranteCoordinado = &elementoXSD{ns: "http://www.sat.gob.mx/pfic", nombre: "PFintegranteCoordinado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "ClaveVehicular", tipo: tPFIC_PFintegranteCoordinado_ClaveVehicular, requerido: true},
		{nombre: "Placa", tipo: tPFIC_PFintegranteCoordinado_Placa, requerido: true},
		{nombre: "RFCPF", tipo: tPFIC_tRFC},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}DocumentoRelacionado
var eGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "DocumentoRelacionado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "OrigenErogacion", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_OrigenErogacion, requerido: true},
		{nombre: "FolioFiscalVinculado", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_FolioFiscalVinculado},
		{nombre: "RFCProveedor", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_RFCProveedor},
		{nombre: "MontoTotalIVA", tipo: tTdCFDI_tImporte},
		{nombre: "MontoRetencionISR", tipo: tTdCFDI_tImporte},
		{nombre: "MontoRetencionIVA", tipo: tTdCFDI_tImporte},
		{nombre: "MontoRetencionOtrosImpuestos", tipo: tTdCFDI_tImporte},
		{nombre: "NumeroPedimentoVinculado", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_NumeroPedimentoVinculado},
		{nombre: "ClavePedimentoVinculado", tipo: tCatHidrocarburos_ClavePedimento},
		{nombre: "ClavePagoPedimentoVinculado", tipo: tCatHidrocarburos_ClavePagoPedimento},
		{nombre: "MontoIVAPedimento", tipo: tTdCFDI_tImporte},
		{nombre: "OtrosImpuestosPagadosPedimento", tipo: tTdCFDI_tImporte},
		{nombre: "FechaFolioFiscalVinculado", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado_FechaFolioFiscalVinculado},
		{nombre: "Mes", tipo: tCatHidrocarburos_Meses, requerido: true},
		{nombre: "MontoTotalErogaciones", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}Tareas
var eGastosHidro_GastosHidrocarburos_Erogacion_Actividades_SubActividades_Tareas = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "Tareas", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TareaRelacionada", tipo: tCatHidrocarburos_Tareas},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}SubActividades
var eGastosHidro_GastosHidrocarburos_Erogacion_Actividades_SubActividades = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "SubActividades", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "SubActividadRelacionada", tipo: tCatHidrocarburos_SubActividad},
	},
	particulas: []particulaXSD{
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_Actividades_SubActividades_Tareas, min: 0, max: -1},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}Actividades
var eGastosHidro_GastosHidrocarburos_Erogacion_Actividades = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "Actividades", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ActividadRelacionada", tipo: tCatHidrocarburos_Actividades},
	},
	particulas: []particulaXSD{
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_Actividades_SubActividades, min: 0, max: -1},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}Pozos
var eGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos_Pozos = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "Pozos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Pozo", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos_Pozos_Pozo},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}Yacimientos
var eGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "Yacimientos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Yacimiento", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos_Yacimiento},
	},
	particulas: []particulaXSD{
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos_Pozos, min: 0, max: -1},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}CentroCostos
var eGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "CentroCostos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Campo", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Campo},
	},
	particulas: []particulaXSD{
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos_Yacimientos, min: 0, max: -1},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}Erogacion
var eGastosHidro_GastosHidrocarburos_Erogacion = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "Erogacion", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "TipoErogacion", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_TipoErogacion, requerido: true},
		{nombre: "MontocuErogacion", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Porcentaje", tipo: tGastosHidro_GastosHidrocarburos_Erogacion_Porcentaje, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_DocumentoRelacionado, min: 1, max: -1},
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_Actividades, min: 0, max: -1},
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion_CentroCostos, min: 0, max: -1},
	}}

// cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd: {http://www.sat.gob.mx/GastosHidrocarburos10}GastosHidrocarburos
var eGastosHidro_GastosHidrocarburos = &elementoXSD{ns: "http://www.sat.gob.mx/GastosHidrocarburos10", nombre: "GastosHidrocarburos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tGastosHidro_GastosHidrocarburos_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "NumeroContrato", tipo: tGastosHidro_GastosHidrocarburos_NumeroContrato, requerido: true},
		{nombre: "AreaContractual", tipo: tGastosHidro_GastosHidrocarburos_AreaContractual},
	},
	particulas: []particulaXSD{
		{elem: eGastosHidro_GastosHidrocarburos_Erogacion, min: 1, max: -1},
	}}

// cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xsd: {http://www.sat.gob.mx/IngresosHidrocarburos10}DocumentoRelacionado
var eIngresosHidro_IngresosHidrocarburos_DocumentoRelacionado = &elementoXSD{ns: "http://www.sat.gob.mx/IngresosHidrocarburos10", nombre: "DocumentoRelacionado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "FolioFiscalVinculado", tipo: tIngresosHidro_IngresosHidrocarburos_DocumentoRelacionado_FolioFiscalVinculado, requerido: true},
		{nombre: "FechaFolioFiscalVinculado", tipo: tIngresosHidro_IngresosHidrocarburos_DocumentoRelacionado_FechaFolioFiscalVinculado, requerido: true},
		{nombre: "Mes", tipo: tCatHidrocarburos_Meses, requerido: true},
	}}

// cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xsd: {http://www.sat.gob.mx/IngresosHidrocarburos10}IngresosHidrocarburos
var eIngresosHidro_IngresosHidrocarburos = &elementoXSD{ns: "http://www.sat.gob.mx/IngresosHidrocarburos10", nombre: "IngresosHidrocarburos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tIngresosHidro_IngresosHidrocarburos_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "NumeroContrato", tipo: tIngresosHidro_IngresosHidrocarburos_NumeroContrato, requerido: true},
		{nombre: "ContraprestacionPagadaOperador", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "Porcentaje", tipo: tIngresosHidro_IngresosHidrocarburos_Porcentaje, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eIngresosHidro_IngresosHidrocarburos_DocumentoRelacionado, min: 1, max: -1},
	}}

// cfd/ecb/ecb.xsd: {http://www.sat.gob.mx/ecb}MovimientoECB
var eECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB = &elementoXSD{ns: "http://www.sat.gob.mx/ecb", nombre: "MovimientoECB", vacio: true,
	attrs: []atributoXSD{
		{nombre: "fecha", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_fecha, requerido: true},
		{nombre: "referencia", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_referencia},
		{nombre: "descripcion", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_descripcion, requerido: true},
		{nombre: "importe", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_importe, requerido: true},
		{nombre: "moneda", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_moneda},
		{nombre: "saldoInicial", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_saldoInicial},
		{nombre: "saldoAlCorte", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB_saldoAlCorte},
	}}

// cfd/ecb/ecb.xsd: {http://www.sat.gob.mx/ecb}MovimientoECBFiscal
var eECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal = &elementoXSD{ns: "http://www.sat.gob.mx/ecb", nombre: "MovimientoECBFiscal", vacio: true,
	attrs: []atributoXSD{
		{nombre: "fecha", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_fecha, requerido: true},
		{nombre: "referencia", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_referencia},
		{nombre: "descripcion", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_descripcion, requerido: true},
		{nombre: "RFCenajenante", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_RFCenajenante, requerido: true},
		{nombre: "Importe", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_Importe, requerido: true},
		{nombre: "moneda", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_moneda},
		{nombre: "saldoInicial", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_saldoInicial},
		{nombre: "saldoAlCorte", tipo: tECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal_saldoAlCorte},
	}}

// cfd/ecb/ecb.xsd: {http://www.sat.gob.mx/ecb}Movimientos
var eECB_EstadoDeCuentaBancario_Movimientos = &elementoXSD{ns: "http://www.sat.gob.mx/ecb", nombre: "Movimientos", seqMin: 1, seqMax: -1,
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eECB_EstadoDeCuentaBancario_Movimientos_MovimientoECB, min: 1, max: 1},
			{elem: eECB_EstadoDeCuentaBancario_Movimientos_MovimientoECBFiscal, min: 1, max: 1},
		}},
	}}

// cfd/ecb/ecb.xsd: {http://www.sat.gob.mx/ecb}EstadoDeCuentaBancario
var eECB_EstadoDeCuentaBancario = &elementoXSD{ns: "http://www.sat.gob.mx/ecb", nombre: "EstadoDeCuentaBancario", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: tECB_EstadoDeCuentaBancario_version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "numeroCuenta", tipo: tECB_EstadoDeCuentaBancario_numeroCuenta, requerido: true},
		{nombre: "nombreCliente", tipo: tECB_EstadoDeCuentaBancario_nombreCliente, requerido: true},
		{nombre: "periodo", tipo: tECB_EstadoDeCuentaBancario_periodo, requerido: true},
		{nombre: "sucursal", tipo: tECB_EstadoDeCuentaBancario_sucursal},
	},
	particulas: []particulaXSD{
		{elem: eECB_EstadoDeCuentaBancario_Movimientos, min: 1, max: 1},
	}}

// cfd/ecc/ecc.xsd: {http://www.sat.gob.mx/ecc}Traslado
var eECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado = &elementoXSD{ns: "http://www.sat.gob.mx/ecc", nombre: "Traslado", vacio: true,
	attrs: []atributoXSD{
		{nombre: "impuesto", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_impuesto, requerido: true},
		{nombre: "tasa", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_tasa, requerido: true},
		{nombre: "importe", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado_importe, requerido: true},
	}}

// cfd/ecc/ecc.xsd: {http://www.sat.gob.mx/ecc}Traslados
var eECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados = &elementoXSD{ns: "http://www.sat.gob.mx/ecc", nombre: "Traslados", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados_Traslado, min: 1, max: -1},
	}}

// cfd/ecc/ecc.xsd: {http://www.sat.gob.mx/ecc}ConceptoEstadoDeCuentaCombustible
var eECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible = &elementoXSD{ns: "http://www.sat.gob.mx/ecc", nombre: "ConceptoEstadoDeCuentaCombustible", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "identificador", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_identificador, requerido: true},
		{nombre: "fecha", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_fecha, requerido: true},
		{nombre: "rfc", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_rfc, requerido: true},
		{nombre: "claveEstacion", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_claveEstacion, requerido: true},
		{nombre: "cantidad", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_cantidad, requerido: true},
		{nombre: "nombreCombustible", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_nombreCombustible, requerido: true},
		{nombre: "folioOperacion", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_folioOperacion, requerido: true},
		{nombre: "valorUnitario", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_valorUnitario, requerido: true},
		{nombre: "importe", tipo: tECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_importe, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible_Traslados, min: 1, max: 1},
	}}

// cfd/ecc/ecc.xsd: {http://www.sat.gob.mx/ecc}Conceptos
var eECC10_EstadoDeCuentaCombustible_Conceptos = &elementoXSD{ns: "http://www.sat.gob.mx/ecc", nombre: "Conceptos", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eECC10_EstadoDeCuentaCombustible_Conceptos_ConceptoEstadoDeCuentaCombustible, min: 1, max: -1},
	}}

// cfd/ecc/ecc.xsd: {http://www.sat.gob.mx/ecc}EstadoDeCuentaCombustible
var eECC10_EstadoDeCuentaCombustible = &elementoXSD{ns: "http://www.sat.gob.mx/ecc", nombre: "EstadoDeCuentaCombustible", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "tipoOperacion", tipo: tECC10_EstadoDeCuentaCombustible_tipoOperacion, requerido: true, fijo: "Tarjeta", conFijo: true},
		{nombre: "numeroDeCuenta", tipo: tECC10_EstadoDeCuentaCombustible_numeroDeCuenta, requerido: true},
		{nombre: "subTotal", tipo: tECC10_EstadoDeCuentaCombustible_subTotal},
		{nombre: "total", tipo: tECC10_EstadoDeCuentaCombustible_total, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eECC10_EstadoDeCuentaCombustible_Conceptos, min: 1, max: 1},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}CfdiRetenRelacionados
var eRet20_Retenciones_CfdiRetenRelacionados = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "CfdiRetenRelacionados", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoRelacion", tipo: tCatCFDI_cTipoRelacion, requerido: true},
		{nombre: "UUID", tipo: tRet20_Retenciones_CfdiRetenRelacionados_UUID, requerido: true},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Emisor
var eRet20_Retenciones_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Emisor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RfcE", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "NomDenRazSocE", tipo: tRet20_Retenciones_Emisor_NomDenRazSocE, requerido: true},
		{nombre: "RegimenFiscalE", tipo: tCatCFDI_cRegimenFiscal, requerido: true},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Nacional
var eRet20_Retenciones_Receptor_Nacional = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Nacional", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RfcR", tipo: tTdCFDI_tRFC, requerido: true},
		{nombre: "NomDenRazSocR", tipo: tRet20_Retenciones_Receptor_Nacional_NomDenRazSocR, requerido: true},
		{nombre: "CurpR", tipo: tTdCFDI_tCURP},
		{nombre: "DomicilioFiscalR", tipo: tRet20_Retenciones_Receptor_Nacional_DomicilioFiscalR, requerido: true},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Extranjero
var eRet20_Retenciones_Receptor_Extranjero = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Extranjero", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTribR", tipo: tRet20_Retenciones_Receptor_Extranjero_NumRegIdTribR},
		{nombre: "NomDenRazSocR", tipo: tRet20_Retenciones_Receptor_Extranjero_NomDenRazSocR, requerido: true},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Receptor
var eRet20_Retenciones_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Receptor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "NacionalidadR", tipo: tRet20_Retenciones_Receptor_NacionalidadR, requerido: true},
	},
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eRet20_Retenciones_Receptor_Nacional, min: 1, max: 1},
			{elem: eRet20_Retenciones_Receptor_Extranjero, min: 1, max: 1},
		}},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Periodo
var eRet20_Retenciones_Periodo = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Periodo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "MesIni", tipo: tCatRetenciones_cPeriodo, requerido: true},
		{nombre: "MesFin", tipo: tCatRetenciones_cPeriodo, requerido: true},
		{nombre: "Ejercicio", tipo: tCatRetenciones_cEjercicio, requerido: true},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}ImpRetenidos
var eRet20_Retenciones_Totales_ImpRetenidos = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "ImpRetenidos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BaseRet", tipo: tTdCFDI_tImporteMXN},
		{nombre: "ImpuestoRet", tipo: tCatCFDI_cImpuesto},
		{nombre: "MontoRet", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "TipoPagoRet", tipo: tCatRetenciones_cTipoPagoRet, requerido: true},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Totales
var eRet20_Retenciones_Totales = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Totales", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "MontoTotOperacion", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "MontoTotGrav", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "MontoTotExent", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "MontoTotRet", tipo: tTdCFDI_tImporteMXN, requerido: true},
		{nombre: "UtilidadBimestral", tipo: tTdCFDI_tImporteMXN},
		{nombre: "ISRCorrespondiente", tipo: tTdCFDI_tImporteMXN},
	},
	particulas: []particulaXSD{
		{elem: eRet20_Retenciones_Totales_ImpRetenidos, min: 0, max: -1},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Complemento
var eRet20_Retenciones_Complemento = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Complemento", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 0, max: -1},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Addenda
var eRet20_Retenciones_Addenda = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Addenda", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 0, max: -1},
	}}

// ../esquemas/retencionpago/2/retencionpagov2.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/2}Retenciones
var eRet20_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/2", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tRet20_Retenciones_Version, requerido: true, fijo: "2.0", conFijo: true},
		{nombre: "FolioInt", tipo: tRet20_Retenciones_FolioInt},
		{nombre: "Sello", tipo: tRet20_Retenciones_Sello, requerido: true},
		{nombre: "NoCertificado", tipo: tRet20_Retenciones_NoCertificado, requerido: true},
		{nombre: "Certificado", tipo: tRet20_Retenciones_Certificado, requerido: true},
		{nombre: "FechaExp", tipo: tTdCFDI_tFechaH, requerido: true},
		{nombre: "LugarExpRetenc", tipo: tCatCFDI_cCodigoPostal, requerido: true},
		{nombre: "CveRetenc", tipo: tCatRetenciones_cCveRetenc, requerido: true},
		{nombre: "DescRetenc", tipo: tRet20_Retenciones_DescRetenc},
	},
	particulas: []particulaXSD{
		{elem: eRet20_Retenciones_CfdiRetenRelacionados, min: 0, max: 1},
		{elem: eRet20_Retenciones_Emisor, min: 1, max: 1},
		{elem: eRet20_Retenciones_Receptor, min: 1, max: 1},
		{elem: eRet20_Retenciones_Periodo, min: 1, max: 1},
		{elem: eRet20_Retenciones_Totales, min: 1, max: 1},
		{elem: eRet20_Retenciones_Complemento, min: 0, max: 1},
		{elem: eRet20_Retenciones_Addenda, min: 0, max: 1},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Emisor
var eRet10_Retenciones_Emisor = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Emisor", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RFCEmisor", tipo: tRet10_tRFC, requerido: true},
		{nombre: "NomDenRazSocE", tipo: tRet10_Retenciones_Emisor_NomDenRazSocE},
		{nombre: "CURPE", tipo: tRet10_tCURP},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Nacional
var eRet10_Retenciones_Receptor_Nacional = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Nacional", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RFCRecep", tipo: tRet10_Retenciones_Receptor_Nacional_RFCRecep, requerido: true},
		{nombre: "NomDenRazSocR", tipo: tRet10_Retenciones_Receptor_Nacional_NomDenRazSocR},
		{nombre: "CURPR", tipo: tRet10_tCURP},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Extranjero
var eRet10_Retenciones_Receptor_Extranjero = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Extranjero", vacio: true,
	attrs: []atributoXSD{
		{nombre: "NumRegIdTrib", tipo: tRet10_Retenciones_Receptor_Extranjero_NumRegIdTrib},
		{nombre: "NomDenRazSocR", tipo: tRet10_Retenciones_Receptor_Extranjero_NomDenRazSocR, requerido: true},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Receptor
var eRet10_Retenciones_Receptor = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Receptor", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Nacionalidad", tipo: tRet10_Retenciones_Receptor_Nacionalidad, requerido: true},
	},
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eRet10_Retenciones_Receptor_Nacional, min: 1, max: 1},
			{elem: eRet10_Retenciones_Receptor_Extranjero, min: 1, max: 1},
		}},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Periodo
var eRet10_Retenciones_Periodo = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Periodo", vacio: true,
	attrs: []atributoXSD{
		{nombre: "MesIni", tipo: tRet10_Retenciones_Periodo_MesIni, requerido: true},
		{nombre: "MesFin", tipo: tRet10_Retenciones_Periodo_MesFin, requerido: true},
		{nombre: "Ejerc", tipo: tRet10_Retenciones_Periodo_Ejerc, requerido: true},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}ImpRetenidos
var eRet10_Retenciones_Totales_ImpRetenidos = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "ImpRetenidos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "BaseRet", tipo: tRet10_tImporte},
		{nombre: "Impuesto", tipo: tRet10_Retenciones_Totales_ImpRetenidos_Impuesto},
		{nombre: "montoRet", tipo: tRet10_tImporte, requerido: true},
		{nombre: "TipoPagoRet", tipo: tRet10_Retenciones_Totales_ImpRetenidos_TipoPagoRet, requerido: true},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Totales
var eRet10_Retenciones_Totales = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Totales", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "montoTotOperacion", tipo: tRet10_tImporte, requerido: true},
		{nombre: "montoTotGrav", tipo: tRet10_tImporte, requerido: true},
		{nombre: "montoTotExent", tipo: tRet10_tImporte, requerido: true},
		{nombre: "montoTotRet", tipo: tRet10_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRet10_Retenciones_Totales_ImpRetenidos, min: 0, max: -1},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Complemento
var eRet10_Retenciones_Complemento = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Complemento", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 0, max: -1},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Addenda
var eRet10_Retenciones_Addenda = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Addenda", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{comodin: true, min: 0, max: -1},
	}}

// ../esquemas/retencionpago/1/retencionpagov1.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1}Retenciones
var eRet10_Retenciones = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1", nombre: "Retenciones", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tRet10_Retenciones_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "FolioInt", tipo: tRet10_Retenciones_FolioInt},
		{nombre: "Sello", tipo: tRet10_Retenciones_Sello, requerido: true},
		{nombre: "NumCert", tipo: tRet10_Retenciones_NumCert, requerido: true},
		{nombre: "Cert", tipo: tRet10_Retenciones_Cert, requerido: true},
		{nombre: "FechaExp", tipo: tRet10_Retenciones_FechaExp, requerido: true},
		{nombre: "CveRetenc", tipo: tCatRetenciones_cRetenciones, requerido: true},
		{nombre: "DescRetenc", tipo: tRet10_Retenciones_DescRetenc},
	},
	particulas: []particulaXSD{
		{elem: eRet10_Retenciones_Emisor, min: 1, max: 1},
		{elem: eRet10_Retenciones_Receptor, min: 1, max: 1},
		{elem: eRet10_Retenciones_Periodo, min: 1, max: 1},
		{elem: eRet10_Retenciones_Totales, min: 1, max: 1},
		{elem: eRet10_Retenciones_Complemento, min: 0, max: 1},
		{elem: eRet10_Retenciones_Addenda, min: 0, max: 1},
	}}

// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10}ImpuestosTrasladadosdelServicio
var eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ImpuestosTrasladadosdelServicio = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", nombre: "ImpuestosTrasladadosdelServicio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ImpuestosTrasladadosdelServicio_Base, requerido: true},
		{nombre: "Impuesto", tipo: tCatRetenciones_cTipoImpuesto, requerido: true},
		{nombre: "TipoFactor", tipo: xsstring, requerido: true, fijo: "Tasa", conFijo: true},
		{nombre: "TasaCuota", tipo: tCatPlataformas_cTasaCuota, requerido: true},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10}ContribucionGubernamental
var eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ContribucionGubernamental = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", nombre: "ContribucionGubernamental", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ImpContrib", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "EntidadDondePagaLaContribucion", tipo: tCatRetenciones_cEntidadesFederativas, requerido: true},
	}}

// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10}ComisionDelServicio
var eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ComisionDelServicio = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", nombre: "ComisionDelServicio", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Base", tipo: tRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ComisionDelServicio_Base},
		{nombre: "Porcentaje", tipo: tRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ComisionDelServicio_Porcentaje},
		{nombre: "Importe", tipo: tTdCFDI_tImporte, requerido: true},
	}}

// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10}DetallesDelServicio
var eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", nombre: "DetallesDelServicio", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "FormaPagoServ", tipo: tCatPlataformas_cFormaPagoServ, requerido: true},
		{nombre: "TipoDeServ", tipo: tCatPlataformas_cTipoDeServ, requerido: true},
		{nombre: "SubTipServ", tipo: tCatPlataformas_cSubTipoServ},
		{nombre: "RFCTerceroAutorizado", tipo: tTdCFDI_tRFCPF},
		{nombre: "FechaServ", tipo: tTdCFDI_tFecha, requerido: true},
		{nombre: "PrecioServSinIVA", tipo: tTdCFDI_tImporte, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ImpuestosTrasladadosdelServicio, min: 0, max: 1},
		{elem: eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ContribucionGubernamental, min: 0, max: 1},
		{elem: eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio_ComisionDelServicio, min: 0, max: 1},
	}}

// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10}Servicios
var eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", nombre: "Servicios", seqMin: 1, seqMax: 1,
	particulas: []particulaXSD{
		{elem: eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios_DetallesDelServicio, min: 1, max: -1},
	}}

// ../esquemas/retencionpago/1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10}ServiciosPlataformasTecnologicas
var eRetPlataformas10_ServiciosPlataformasTecnologicas = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10", nombre: "ServiciosPlataformasTecnologicas", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "Periodicidad", tipo: tCatRetenciones_cPeriodicidad, requerido: true},
		{nombre: "NumServ", tipo: tRetPlataformas10_ServiciosPlataformasTecnologicas_NumServ, requerido: true},
		{nombre: "MonTotServSIVA", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "TotalIVATrasladado", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "TotalIVARetenido", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "TotalISRRetenido", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "DifIVAEntregadoPrestServ", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "MonTotalporUsoPlataforma", tipo: tTdCFDI_tImporte, requerido: true},
		{nombre: "MonTotalContribucionGubernamental", tipo: tTdCFDI_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eRetPlataformas10_ServiciosPlataformasTecnologicas_Servicios, min: 1, max: 1},
	}}

// ../esquemas/retencionpago/1/intereses/intereses.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/intereses}Intereses
var eRetIntereses_Intereses = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/intereses", nombre: "Intereses", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "SistFinanciero", tipo: tRetIntereses_Intereses_SistFinanciero, requerido: true},
		{nombre: "RetiroAORESRetInt", tipo: tRetIntereses_Intereses_RetiroAORESRetInt, requerido: true},
		{nombre: "OperFinancDerivad", tipo: tRetIntereses_Intereses_OperFinancDerivad, requerido: true},
		{nombre: "MontIntNominal", tipo: tRetIntereses_tImporte, requerido: true},
		{nombre: "MontIntReal", tipo: tRetIntereses_tImporte, requerido: true},
		{nombre: "Perdida", tipo: tRetIntereses_tImporte, requerido: true},
	}}

// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros}NoBeneficiario
var eRetPagosExtranjeros_Pagosaextranjeros_NoBeneficiario = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", nombre: "NoBeneficiario", vacio: true,
	attrs: []atributoXSD{
		{nombre: "PaisDeResidParaEfecFisc", tipo: tCatRetenciones_cPais, requerido: true},
		{nombre: "ConceptoPago", tipo: tCatRetenciones_cTipoContribuyenteSujetoRetencion, requerido: true},
		{nombre: "DescripcionConcepto", tipo: tRetPagosExtranjeros_Pagosaextranjeros_NoBeneficiario_DescripcionConcepto, requerido: true},
	}}

// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros}Beneficiario
var eRetPagosExtranjeros_Pagosaextranjeros_Beneficiario = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", nombre: "Beneficiario", vacio: true,
	attrs: []atributoXSD{
		{nombre: "RFC", tipo: tRetPagosExtranjeros_tRFC, requerido: true},
		{nombre: "CURP", tipo: tRetPagosExtranjeros_tCURP, requerido: true},
		{nombre: "NomDenRazSocB", tipo: tRetPagosExtranjeros_Pagosaextranjeros_Beneficiario_NomDenRazSocB, requerido: true},
		{nombre: "ConceptoPago", tipo: tCatRetenciones_cTipoContribuyenteSujetoRetencion, requerido: true},
		{nombre: "DescripcionConcepto", tipo: tRetPagosExtranjeros_Pagosaextranjeros_Beneficiario_DescripcionConcepto, requerido: true},
	}}

// ../esquemas/retencionpago/1/pagosaextranjeros/pagosaextranjeros.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros}Pagosaextranjeros
var eRetPagosExtranjeros_Pagosaextranjeros = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros", nombre: "Pagosaextranjeros", seqMin: 0, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "EsBenefEfectDelCobro", tipo: tRetPagosExtranjeros_Pagosaextranjeros_EsBenefEfectDelCobro, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRetPagosExtranjeros_Pagosaextranjeros_NoBeneficiario, min: 0, max: 1},
		{elem: eRetPagosExtranjeros_Pagosaextranjeros_Beneficiario, min: 0, max: 1},
	}}

// ../esquemas/retencionpago/1/dividendos/dividendos.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos}DividOUtil
var eRetDividendos_Dividendos_DividOUtil = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", nombre: "DividOUtil", vacio: true,
	attrs: []atributoXSD{
		{nombre: "CveTipDivOUtil", tipo: tCatRetenciones_cTipoDividendoOUtilidadDistribuida, requerido: true},
		{nombre: "MontISRAcredRetMexico", tipo: tRetDividendos_tImporte, requerido: true},
		{nombre: "MontISRAcredRetExtranjero", tipo: tRetDividendos_tImporte, requerido: true},
		{nombre: "MontRetExtDivExt", tipo: tRetDividendos_tImporte},
		{nombre: "TipoSocDistrDiv", tipo: tRetDividendos_Dividendos_DividOUtil_TipoSocDistrDiv, requerido: true},
		{nombre: "MontISRAcredNal", tipo: tRetDividendos_tImporte},
		{nombre: "MontDivAcumNal", tipo: tRetDividendos_tImporte},
		{nombre: "MontDivAcumExt", tipo: tRetDividendos_tImporte},
	}}

// ../esquemas/retencionpago/1/dividendos/dividendos.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos}Remanente
var eRetDividendos_Dividendos_Remanente = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", nombre: "Remanente", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ProporcionRem", tipo: tRetDividendos_Dividendos_Remanente_ProporcionRem},
	}}

// ../esquemas/retencionpago/1/dividendos/dividendos.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos}Dividendos
var eRetDividendos_Dividendos = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos", nombre: "Dividendos", seqMin: 0, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
	},
	particulas: []particulaXSD{
		{elem: eRetDividendos_Dividendos_DividOUtil, min: 0, max: 1},
		{elem: eRetDividendos_Dividendos_Remanente, min: 0, max: 1},
	}}

// ../esquemas/retencionpago/1/enajenaciondeacciones/enajenaciondeacciones.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones}EnajenaciondeAcciones
var eRetEnajenacion_EnajenaciondeAcciones = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones", nombre: "EnajenaciondeAcciones", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "ContratoIntermediacion", tipo: tRetEnajenacion_EnajenaciondeAcciones_ContratoIntermediacion, requerido: true},
		{nombre: "Ganancia", tipo: tRetEnajenacion_tImporte, requerido: true},
		{nombre: "Perdida", tipo: tRetEnajenacion_tImporte, requerido: true},
	}}

// ../esquemas/retencionpago/1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso}Arrendamientoenfideicomiso
var eRetArrendamiento_Arrendamientoenfideicomiso = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso", nombre: "Arrendamientoenfideicomiso", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "PagProvEfecPorFiduc", tipo: tRetArrendamiento_tImporte, requerido: true},
		{nombre: "RendimFideicom", tipo: tRetArrendamiento_tImporte, requerido: true},
		{nombre: "DeduccCorresp", tipo: tRetArrendamiento_tImporte, requerido: true},
		{nombre: "MontTotRet", tipo: tRetArrendamiento_tImporte},
		{nombre: "MontResFiscDistFibras", tipo: tRetArrendamiento_tImporte},
		{nombre: "MontOtrosConceptDistr", tipo: tRetArrendamiento_tImporte},
		{nombre: "DescrMontOtrosConceptDistr", tipo: xsstring},
	}}

// ../esquemas/retencionpago/1/operacionesconderivados/operacionesconderivados.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados}Operacionesconderivados
var eRetDerivados_Operacionesconderivados = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados", nombre: "Operacionesconderivados", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "MontGanAcum", tipo: tRetDerivados_Operacionesconderivados_MontGanAcum, requerido: true},
		{nombre: "MontPerdDed", tipo: tRetDerivados_tImporte, requerido: true},
	}}

// ../esquemas/retencionpago/1/intereseshipotecarios/intereseshipotecarios.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios}Intereseshipotecarios
var eRetHipotecarios_Intereseshipotecarios = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios", nombre: "Intereseshipotecarios", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "CreditoDeInstFinanc", tipo: tRetHipotecarios_Intereseshipotecarios_CreditoDeInstFinanc, requerido: true},
		{nombre: "SaldoInsoluto", tipo: tRetHipotecarios_tImporte, requerido: true},
		{nombre: "PropDeducDelCredit", tipo: tRetHipotecarios_tImporte},
		{nombre: "MontTotIntNominalesDev", tipo: tRetHipotecarios_tImporte},
		{nombre: "MontTotIntNominalesDevYPag", tipo: tRetHipotecarios_tImporte},
		{nombre: "MontTotIntRealPagDeduc", tipo: tRetHipotecarios_tImporte},
		{nombre: "NumContrato", tipo: tRetHipotecarios_Intereseshipotecarios_NumContrato},
	}}

// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial}IntegracIngresos
var eRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas_IntegracIngresos = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", nombre: "IntegracIngresos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Concepto", tipo: tRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas_IntegracIngresos_Concepto, requerido: true},
	}}

// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial}IngresosOEntradas
var eRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", nombre: "IngresosOEntradas", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "MontTotEntradasPeriodo", tipo: tRetFideicomiso_tImporte, requerido: true},
		{nombre: "PartPropAcumDelFideicom", tipo: tRetFideicomiso_tImporte, requerido: true},
		{nombre: "PropDelMontTot", tipo: tRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas_PropDelMontTot, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas_IntegracIngresos, min: 1, max: 1},
	}}

// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial}IntegracEgresos
var eRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas_IntegracEgresos = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", nombre: "IntegracEgresos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "ConceptoS", tipo: tRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas_IntegracEgresos_ConceptoS, requerido: true},
	}}

// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial}DeduccOSalidas
var eRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", nombre: "DeduccOSalidas", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "MontTotEgresPeriodo", tipo: tRetFideicomiso_tImporte, requerido: true},
		{nombre: "PartPropDelFideicom", tipo: tRetFideicomiso_tImporte, requerido: true},
		{nombre: "PropDelMontTot", tipo: tRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas_PropDelMontTot, requerido: true},
	},
	particulas: []particulaXSD{
		{elem: eRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas_IntegracEgresos, min: 1, max: 1},
	}}

// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial}RetEfectFideicomiso
var eRetFideicomiso_Fideicomisonoempresarial_RetEfectFideicomiso = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", nombre: "RetEfectFideicomiso", vacio: true,
	attrs: []atributoXSD{
		{nombre: "MontRetRelPagFideic", tipo: tRetFideicomiso_tImporte, requerido: true},
		{nombre: "DescRetRelPagFideic", tipo: tRetFideicomiso_Fideicomisonoempresarial_RetEfectFideicomiso_DescRetRelPagFideic, requerido: true},
	}}

// ../esquemas/retencionpago/1/fideicomisonoempresarial/fideicomisonoempresarial.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial}Fideicomisonoempresarial
var eRetFideicomiso_Fideicomisonoempresarial = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial", nombre: "Fideicomisonoempresarial", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
	},
	particulas: []particulaXSD{
		{elem: eRetFideicomiso_Fideicomisonoempresarial_IngresosOEntradas, min: 1, max: 1},
		{elem: eRetFideicomiso_Fideicomisonoempresarial_DeduccOSalidas, min: 1, max: 1},
		{elem: eRetFideicomiso_Fideicomisonoempresarial_RetEfectFideicomiso, min: 1, max: 1},
	}}

// ../esquemas/retencionpago/1/planesderetiro/planesderetiro.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro}Planesderetiro
var eRetPlanesRetiro10_Planesderetiro = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro", nombre: "Planesderetiro", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "SistemaFinanc", tipo: tRetPlanesRetiro10_Planesderetiro_SistemaFinanc, requerido: true},
		{nombre: "MontTotAportAnioInmAnterior", tipo: tRetPlanesRetiro10_tImporte},
		{nombre: "MontIntRealesDevengAniooInmAnt", tipo: tRetPlanesRetiro10_tImporte, requerido: true},
		{nombre: "HuboRetirosAnioInmAntPer", tipo: tRetPlanesRetiro10_Planesderetiro_HuboRetirosAnioInmAntPer, requerido: true},
		{nombre: "MontTotRetiradoAnioInmAntPer", tipo: tRetPlanesRetiro10_tImporte},
		{nombre: "MontTotExentRetiradoAnioInmAnt", tipo: tRetPlanesRetiro10_tImporte},
		{nombre: "MontTotExedenteAnioInmAnt", tipo: tRetPlanesRetiro10_tImporte},
		{nombre: "HuboRetirosAnioInmAnt", tipo: tRetPlanesRetiro10_Planesderetiro_HuboRetirosAnioInmAnt, requerido: true},
		{nombre: "MontTotRetiradoAnioInmAnt", tipo: tRetPlanesRetiro10_tImporte},
	}}

// ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11}AportacionesODepositos
var eRetPlanesRetiro11_Planesderetiro_AportacionesODepositos = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", nombre: "AportacionesODepositos", vacio: true,
	attrs: []atributoXSD{
		{nombre: "TipoAportacionODeposito", tipo: tCatPlanesRetiro_cTipoAportODep, requerido: true},
		{nombre: "MontAportODep", tipo: tRetPlanesRetiro11_tImporte, requerido: true},
		{nombre: "RFCFiduciaria", tipo: tRetPlanesRetiro11_tRFC},
	}}

// ../esquemas/retencionpago/1/planesderetiro11/planesderetiro11.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11}Planesderetiro
var eRetPlanesRetiro11_Planesderetiro = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11", nombre: "Planesderetiro", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.1", conFijo: true},
		{nombre: "SistemaFinanc", tipo: tRetPlanesRetiro11_Planesderetiro_SistemaFinanc, requerido: true},
		{nombre: "MontTotAportAnioInmAnterior", tipo: tRetPlanesRetiro11_tImporte},
		{nombre: "MontIntRealesDevengAniooInmAnt", tipo: tRetPlanesRetiro11_tImporte, requerido: true},
		{nombre: "HuboRetirosAnioInmAntPer", tipo: tRetPlanesRetiro11_Planesderetiro_HuboRetirosAnioInmAntPer, requerido: true},
		{nombre: "MontTotRetiradoAnioInmAntPer", tipo: tRetPlanesRetiro11_tImporte},
		{nombre: "MontTotExentRetiradoAnioInmAnt", tipo: tRetPlanesRetiro11_tImporte},
		{nombre: "MontTotExedenteAnioInmAnt", tipo: tRetPlanesRetiro11_tImporte},
		{nombre: "HuboRetirosAnioInmAnt", tipo: tRetPlanesRetiro11_Planesderetiro_HuboRetirosAnioInmAnt, requerido: true},
		{nombre: "MontTotRetiradoAnioInmAnt", tipo: tRetPlanesRetiro11_tImporte},
		{nombre: "NumReferencia", tipo: tRetPlanesRetiro11_Planesderetiro_NumReferencia},
	},
	particulas: []particulaXSD{
		{elem: eRetPlanesRetiro11_Planesderetiro_AportacionesODepositos, min: 0, max: 3},
	}}

// ../esquemas/retencionpago/1/premios/premios.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/premios}Premios
var eRetPremios_Premios = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/premios", nombre: "Premios", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "EntidadFederativa", tipo: tCatRetenciones_cEntidadesFederativas, requerido: true},
		{nombre: "MontTotPago", tipo: tRetPremios_tImporte, requerido: true},
		{nombre: "MontTotPagoGrav", tipo: tRetPremios_tImporte, requerido: true},
		{nombre: "MontTotPagoExent", tipo: tRetPremios_tImporte, requerido: true},
	}}

// ../esquemas/retencionpago/1/sectorfinanciero/sectorfinanciero.xsd: {http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero}SectorFinanciero
var eRetSectorFinanciero_SectorFinanciero = &elementoXSD{ns: "http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero", nombre: "SectorFinanciero", vacio: true,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "IdFideicom", tipo: tRetSectorFinanciero_SectorFinanciero_IdFideicom, requerido: true},
		{nombre: "NomFideicom", tipo: tRetSectorFinanciero_SectorFinanciero_NomFideicom},
		{nombre: "DescripFideicom", tipo: tRetSectorFinanciero_SectorFinanciero_DescripFideicom, requerido: true},
	}}

// cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd: {http://www.sat.gob.mx/TimbreFiscalDigital}TimbreFiscalDigital
var eTFD10_TimbreFiscalDigital = &elementoXSD{ns: "http://www.sat.gob.mx/TimbreFiscalDigital", nombre: "TimbreFiscalDigital", vacio: true,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "UUID", tipo: tTFD10_TimbreFiscalDigital_UUID, requerido: true},
		{nombre: "FechaTimbrado", tipo: tTFD10_TimbreFiscalDigital_FechaTimbrado, requerido: true},
		{nombre: "selloCFD", tipo: tTFD10_TimbreFiscalDigital_selloCFD, requerido: true},
		{nombre: "noCertificadoSAT", tipo: tTFD10_TimbreFiscalDigital_noCertificadoSAT, requerido: true},
		{nombre: "selloSAT", tipo: tTFD10_TimbreFiscalDigital_selloSAT, requerido: true},
	}}

// cfd/donat/donat.xsd: {http://www.sat.gob.mx/donat}Donatarias
var eDonat10_Donatarias = &elementoXSD{ns: "http://www.sat.gob.mx/donat", nombre: "Donatarias", vacio: true,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsstring, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "noAutorizacion", tipo: xsstring, requerido: true},
		{nombre: "fechaAutorizacion", tipo: xsdate, requerido: true},
		{nombre: "leyenda", tipo: tDonat10_Donatarias_leyenda, requerido: true, fijo: "Este comprobante ampara un donativo, el cual será destinado por la donataria a los fines propios de su objeto social. En el caso de que los bienes donados hayan sido deducidos previamente para los efectos del impuesto sobre la renta, este donativo no es deducible. La reproducción no autorizada de este comprobante constituye un delito en los términos de las disposiciones fiscales.", conFijo: true},
	}}

// cfd/ventavehiculos/ventavehiculos.xsd: {http://www.sat.gob.mx/ventavehiculos}InformacionAduanera
var eVentaVehiculos10_VentaVehiculos_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tVentaVehiculos10_VentaVehiculos_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tVentaVehiculos10_VentaVehiculos_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tVentaVehiculos10_VentaVehiculos_InformacionAduanera_aduana},
	}}

// cfd/ventavehiculos/ventavehiculos.xsd: {http://www.sat.gob.mx/ventavehiculos}InformacionAduanera
var eVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "InformacionAduanera", vacio: true,
	attrs: []atributoXSD{
		{nombre: "numero", tipo: tVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera_numero, requerido: true},
		{nombre: "fecha", tipo: tVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera_fecha, requerido: true},
		{nombre: "aduana", tipo: tVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera_aduana},
	}}

// cfd/ventavehiculos/ventavehiculos.xsd: {http://www.sat.gob.mx/ventavehiculos}Parte
var eVentaVehiculos10_VentaVehiculos_Parte = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "Parte", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "cantidad", tipo: tVentaVehiculos10_VentaVehiculos_Parte_cantidad, requerido: true},
		{nombre: "unidad", tipo: tVentaVehiculos10_VentaVehiculos_Parte_unidad},
		{nombre: "noIdentificacion", tipo: tVentaVehiculos10_VentaVehiculos_Parte_noIdentificacion},
		{nombre: "descripcion", tipo: tVentaVehiculos10_VentaVehiculos_Parte_descripcion, requerido: true},
		{nombre: "valorUnitario", tipo: tVentaVehiculos10_tImporte},
		{nombre: "importe", tipo: tVentaVehiculos10_tImporte},
	},
	particulas: []particulaXSD{
		{elem: eVentaVehiculos10_VentaVehiculos_Parte_InformacionAduanera, min: 0, max: -1},
	}}

// cfd/ventavehiculos/ventavehiculos.xsd: {http://www.sat.gob.mx/ventavehiculos}VentaVehiculos
var eVentaVehiculos10_VentaVehiculos = &elementoXSD{ns: "http://www.sat.gob.mx/ventavehiculos", nombre: "VentaVehiculos", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "version", tipo: xsanySimpleType, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "ClaveVehicular", tipo: tVentaVehiculos10_VentaVehiculos_ClaveVehicular, requerido: true},
	},
	particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{
			{elem: eVentaVehiculos10_VentaVehiculos_InformacionAduanera, min: 0, max: -1},
			{elem: eVentaVehiculos10_VentaVehiculos_Parte, min: 0, max: -1},
		}},
	}}

// cfd/ine/ine10.xsd: {http://www.sat.gob.mx/ine}Contabilidad
var eINE10_INE_Entidad_Contabilidad = &elementoXSD{ns: "http://www.sat.gob.mx/ine", nombre: "Contabilidad", vacio: true,
	attrs: []atributoXSD{
		{nombre: "IdContabilidad", tipo: tINE10_INE_Entidad_Contabilidad_IdContabilidad, requerido: true},
	}}

// cfd/ine/ine10.xsd: {http://www.sat.gob.mx/ine}Entidad
var eINE10_INE_Entidad = &elementoXSD{ns: "http://www.sat.gob.mx/ine", nombre: "Entidad", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "ClaveEntidad", tipo: tINE10_tClaveEntidad, requerido: true},
		{nombre: "Ambito", tipo: tINE10_INE_Entidad_Ambito},
	},
	particulas: []particulaXSD{
		{elem: eINE10_INE_Entidad_Contabilidad, min: 0, max: -1},
	}}

// cfd/ine/ine10.xsd: {http://www.sat.gob.mx/ine}INE
var eINE10_INE = &elementoXSD{ns: "http://www.sat.gob.mx/ine", nombre: "INE", seqMin: 1, seqMax: 1,
	attrs: []atributoXSD{
		{nombre: "Version", tipo: tINE10_INE_Version, requerido: true, fijo: "1.0", conFijo: true},
		{nombre: "TipoProceso", tipo: tINE10_INE_TipoProceso, requerido: true},
		{nombre: "TipoComite", tipo: tINE10_INE_TipoComite},
		{nombre: "IdContabilidad", tipo: tINE10_INE_IdContabilidad},
	},
	particulas: []particulaXSD{
		{elem: eINE10_INE_Entidad, min: 0, max: 64},
	}}
