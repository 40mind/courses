package yookassa_provider

import (
    "bytes"
    "courses/domain/models"
    "encoding/json"
    "fmt"
    "github.com/jimlawless/whereami"
    "io"
    "log"
    "net/http"
    "time"
)

const yookassaProviderError = "yookassa provider error"

type Provider struct {
    createPaymentBuilder        endpointBuilder
    getPaymentBuilder           endpointBuilder
    auth                        models.Auth
}

type endpointBuilder struct {
    host            string
    path            string
    method          string
}

func NewProvider(config models.Provider, auth models.Auth) Provider {
    createPaymentBuilder := endpointBuilder{
        host:   config.Host,
        path:   config.Endpoint["CreatePayment"].Path,
        method: config.Endpoint["CreatePayment"].Method,
    }

    getPaymentBuilder := endpointBuilder{
        host:   config.Host,
        path:   config.Endpoint["GetPayment"].Path,
        method: config.Endpoint["GetPayment"].Method,
    }

    return Provider{
        createPaymentBuilder: createPaymentBuilder,
        getPaymentBuilder:    getPaymentBuilder,
        auth:                 auth,
    }
}

func (p Provider) CreatePayment(body []byte, idempotenceKey string) (models.GetPayment, error) {
    client := http.Client{Timeout: 5 * time.Second}

    req, err := p.createPaymentBuilder.createRequest(body, "")
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    h := make(http.Header)
    h.Add("Content-Type", "application/json")
    h.Add("Idempotence-Key", idempotenceKey)
    req.Header = h
    req.SetBasicAuth(p.auth.Login, p.auth.Password)

    response, err := client.Do(req)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    respBody, err := io.ReadAll(response.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    if response.StatusCode != http.StatusOK {
        log.Printf("%s: status code %d: body %s: %s\n", yookassaProviderError, response.StatusCode, string(respBody), whereami.WhereAmI())
        return models.GetPayment{}, fmt.Errorf("%s: status code %d: body %s", yookassaProviderError, response.StatusCode, string(respBody))
    }

    var payment models.GetPayment
    err = json.Unmarshal(respBody, &payment)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    return payment, nil
}

func (p Provider) GetPayment(id string) (models.GetPayment, error) {
    client := http.Client{Timeout: 5 * time.Second}

    req, err := p.getPaymentBuilder.createRequest(nil, "/" + id)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }
    req.SetBasicAuth(p.auth.Login, p.auth.Password)

    response, err := client.Do(req)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    respBody, err := io.ReadAll(response.Body)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    if response.StatusCode != http.StatusOK {
        log.Printf("%s: status code %d: body %s: %s\n", yookassaProviderError, response.StatusCode, string(respBody), whereami.WhereAmI())
        return models.GetPayment{}, fmt.Errorf("%s: status code %d: body %s", yookassaProviderError, response.StatusCode, string(respBody))
    }

    var payment models.GetPayment
    err = json.Unmarshal(respBody, &payment)
    if err != nil {
        log.Printf("%s: %s: %s\n", yookassaProviderError, err.Error(), whereami.WhereAmI())
        return models.GetPayment{}, err
    }

    return payment, nil
}

func (e endpointBuilder) createRequest(body []byte, path string) (*http.Request, error) {
    return http.NewRequest(
        e.method,
        e.host + e.path + path,
        bytes.NewReader(body),
    )
}
