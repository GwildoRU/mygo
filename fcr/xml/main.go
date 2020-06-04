package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Individual struct {
		Surname    string `xml:"surname"`
		Name       string `xml:"name"`
		Patronymic string `xml:"patronymic"`
		BirthDate  string `xml:"birth_date"`
		BirthPlace string `xml:"birth_place"`
	}

	type RightHolder struct {
		I Individual `xml:"individual"`
	}

	type RightHolders struct {
		Rh RightHolder `xml:"right_holder"`
	}

	type RightRecord struct {
		Rhs RightHolders `xml:"right_holders"`
	}

	type RightRecords struct {
		Rr []RightRecord `xml:"right_record"`
	}



	type Base struct {
		XMLName xml.Name `xml:"extract_base_params_room"`
		Rrs  RightRecords `xml:"right_records"`
	}

	v := Base{}


//	<?xml version="1.0" encoding="WINDOWS-1251"?>

	data := `
<extract_base_params_room>
    <details_statement>
        <group_top_requisites>
            <organ_registr_rights>Филиал Федерального государственного бюджетного учреждения &quot;Федеральная
                кадастровая палата Федеральной службы государственной регистрации, кадастра и картографии&quot; по
                Тамбовской области
            </organ_registr_rights>
            <date_formation>2020-03-17</date_formation>
            <registration_number>КУВИ-001/2020-5779088</registration_number>
        </group_top_requisites>
    </details_statement>
    <details_request>
        <date_received_request>2020-03-17</date_received_request>
        <date_receipt_request_reg_authority_rights>2020-03-18</date_receipt_request_reg_authority_rights>
    </details_request>
    <room_record>
        <record_info>
            <registration_date>2012-06-30T00:00:00+04:00</registration_date>
        </record_info>
        <object>
            <common_data>
                <cad_number>68:29:0101048:186</cad_number>
                <quarter_cad_number>68:29:0101048</quarter_cad_number>
                <type>
                    <code>002001003000</code>
                    <value>room_record</value>
                </type>
            </common_data>
        </object>
        <cad_links>
            <parent_cad_number>
                <cad_number>68:29:0101048:51</cad_number>
            </parent_cad_number>
        </cad_links>
        <params>
            <area>137.2</area>
            <name>Торговое помещение</name>
            <purpose>
                <code>206001000000</code>
                <value>Нежилое</value>
            </purpose>
            <common_property>false</common_property>
            <service>false</service>
        </params>
        <address_room>
            <address>
                <address>
                    <readable_address>Тамбовская обл., г.Тамбов, ул.Советская/Ленинградская, д.88/14, пом.1
                    </readable_address>
                </address>
            </address>
        </address_room>
        <location_in_build>
            <level>
                <floor>1</floor>
            </level>
        </location_in_build>
        <cost>
            <value>7490513.58</value>
        </cost>
    </room_record>
    <right_records>
        <right_record>
            <record_info>
                <registration_date>2018-04-23T10:40:55+03:00</registration_date>
            </record_info>
            <right_data>
                <right_type>
                    <code>001002000000</code>
                    <value>Общая долевая собственность</value>
                </right_type>
                <right_number>68:29:0101048:186-68/001/2018-8</right_number>
                <shares>
                    <share>
                        <numerator>1</numerator>
                        <denominator>3</denominator>
                    </share>
                </shares>
            </right_data>
            <right_holders>
                <right_holder>
                    <individual>
                        <surname>Смирнов</surname>
                        <name>Алексей</name>
                        <patronymic>Александрович</patronymic>
                        <birth_date>1986-05-22</birth_date>
                        <birth_place>п. Черноморское Крымской области</birth_place>
                        <citizenship>
                            <person_citizenship_country>
                                <citizenship_country>
                                    <code>848000000643</code>
                                    <value>Российская Федерация</value>
                                </citizenship_country>
                            </person_citizenship_country>
                        </citizenship>
                        <snils>128-142-547 44</snils>
                        <identity_doc>
                            <document_code>
                                <code>008001001000</code>
                                <value>Паспорт гражданина Российской Федерации</value>
                            </document_code>
                            <document_name>паспорт гражданина Российской Федерации</document_name>
                            <document_series>63 05</document_series>
                            <document_number>809692</document_number>
                            <document_date>2006-07-24</document_date>
                            <document_issuer>Отделом внутренних дел Турковского района Саратовской области
                            </document_issuer>
                        </identity_doc>
                        <contacts>
                            <mailing_addess>Московская обл., Одинцовский район, пгт. Заречье, ул. Березовая, дом №9, кв.
                                13
                            </mailing_addess>
                        </contacts>
                    </individual>
                </right_holder>
            </right_holders>
        </right_record>
        <right_record>
            <record_info>
                <registration_date>2017-09-19T22:49:07+03:00</registration_date>
            </record_info>
            <right_data>
                <right_type>
                    <code>001002000000</code>
                    <value>Общая долевая собственность</value>
                </right_type>
                <right_number>68:29:0101048:186-68/001/2017-6</right_number>
                <shares>
                    <share>
                        <numerator>2</numerator>
                        <denominator>3</denominator>
                    </share>
                </shares>
            </right_data>
            <right_holders>
                <right_holder>
                    <individual>
                        <surname>Нуйкина</surname>
                        <name>Татьяна</name>
                        <patronymic>Анатольевна</patronymic>
                        <birth_date>1988-04-27</birth_date>
                        <birth_place>гор. Карловы Вары ЧССР</birth_place>
                        <citizenship>
                            <person_citizenship_country>
                                <citizenship_country>
                                    <code>848000000643</code>
                                    <value>Российская Федерация</value>
                                </citizenship_country>
                            </person_citizenship_country>
                        </citizenship>
                        <snils>128-390-406 61</snils>
                        <identity_doc>
                            <document_code>
                                <code>008001001000</code>
                                <value>Паспорт гражданина Российской Федерации</value>
                            </document_code>
                            <document_name>паспорт гражданина Российской Федерации</document_name>
                            <document_series>68 12</document_series>
                            <document_number>724377</document_number>
                            <document_date>2012-08-21</document_date>
                            <document_issuer>Отдел УФМС России по Тамбовской области в Советском р-не г.Тамбова
                            </document_issuer>
                        </identity_doc>
                        <contacts>
                            <mailing_addess>Тамбовская обл., г. Тамбов, ул. Бастионная, дом №24А, кв. 267
                            </mailing_addess>
                        </contacts>
                    </individual>
                </right_holder>
            </right_holders>
        </right_record>
    </right_records>
    <restrict_records>
        <restrict_record>
            <record_info>
                <registration_date>2010-11-03T00:00:00+03:00</registration_date>
            </record_info>
            <restrictions_encumbrances_data>
                <restriction_encumbrance_number>68-68-01/057/2010-978</restriction_encumbrance_number>
                <restriction_encumbrance_type>
                    <code>022015000000</code>
                    <value>Объект культурного наследия</value>
                </restriction_encumbrance_type>
                <restricting_rights>
                    <restricting_right>
                        <number>68:29:0101048:186-68.egrppart.J.6801:1131463</number>
                        <right_number>68:29:0101048:186-68/001/2018-8</right_number>
                    </restricting_right>
                    <restricting_right>
                        <number>68:29:0101048:186-68.egrppart.J.6801:1055512</number>
                        <right_number>68:29:0101048:186-68/001/2017-6</right_number>
                    </restricting_right>
                </restricting_rights>
            </restrictions_encumbrances_data>
            <right_holders>
                <right_holder>
                    <legal_entity>
                        <entity>
                            <resident>
                                <name>Управление культуры и архивного дела Тамбовской области</name>
                                <inn>6829019043</inn>
                            </resident>
                        </entity>
                    </legal_entity>
                </right_holder>
            </right_holders>
            <underlying_documents>
                <underlying_document>
                    <document_code>
                        <code>558401010101</code>
                        <value>Договор купли-продажи</value>
                    </document_code>
                    <document_name>Договор купли-продажи арендованного муниципального имущества</document_name>
                    <document_number>№ 1162</document_number>
                    <document_date>2009-12-16</document_date>
                </underlying_document>
                <underlying_document>
                    <document_code>
                        <value>Прочий документ</value>
                    </document_code>
                    <document_name>охранное обязательство по использованию объекта культурного наследия</document_name>
                    <document_date>2008-01-24</document_date>
                </underlying_document>
                <underlying_document>
                    <document_code>
                        <value>Прочий документ</value>
                    </document_code>
                    <document_name>Соглашение об отступном</document_name>
                    <document_date>2017-09-06</document_date>
                </underlying_document>
                <underlying_document>
                    <document_code>
                        <value>(Договор общий документ)</value>
                    </document_code>
                    <document_name>Соглашение об отступном</document_name>
                    <document_date>2018-04-13</document_date>
                </underlying_document>
            </underlying_documents>
        </restrict_record>
    </restrict_records>
    <recipient_statement>Управление государственного жилищного надзора Тамбовской области</recipient_statement>
    <status>Сведения об объекте недвижимости имеют статус &quot;актуальные, ранее учтенные&quot;</status>
</extract_base_params_room>	
`
	xml.Unmarshal([]byte(data), &v)

	fmt.Printf("%#v\n", v)

}
